package TCP

import (
	"bufio"
	"fmt"
	"go_cache_server/cache"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	cache.Cache
}

func New(c cache.Cache) *Server  {
	return &Server{c}
}

func (s *Server) Listen() {
	l, e := net.Listen("tcp", "8888")
	if e != nil {
		panic(e)
	}
	for true {
		c, e := l.Accept()
		if e != nil {
			panic(e)
		}
		go s.process(c)
	}
}

func (s *Server) process(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for true {
		op, e := r.ReadByte()
		if e != nil {
			if e != io.EOF {
				log.Println("close connection due to error:", e)
				return
			}
		}
		if op == 's' {
			e = s.set(conn,r)
		}else if op == 'g' {
			e = s.get(conn,r)
		}else if op == 'd' {
			e = s.del(conn,r)
		}else {
			log.Println("close connection due to invalid operation:",op)
			return
		}
		if e != nil {
			log.Println("close connection due to error:",e)
			return
		}
	}
}

func readLen(r *bufio.Reader) (int, error) {
	tmp, e := r.ReadString(' ')
	if e != nil {
		return 0, e
	}
	//删除尾随空格
	l, e := strconv.Atoi(strings.TrimSpace(tmp))
	if e != nil {
		return 0, e
	}
	return l, nil
}

//解析协议里的key，用于get，del
func (s *Server) readKey(r *bufio.Reader) (string, error) {
	klen, e := readLen(r)
	if e != nil {
		return "", e
	}
	k := make([]byte, klen)
	_, e = io.ReadFull(r, k)
	if e != nil {
		return "", e
	}
	return string(k), nil
}

//解析协议里的key，value，用于set
func (s *Server) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	klen, e := readLen(r)
	if e != nil {
		return "", nil, e
	}
	vlen, e := readLen(r)
	if e != nil {
		return "", nil, e
	}

	/*
		ReadFull将r中的len（buf）个字节准确地读取到buf中。它返回复制的字节数，如果读取的字节数少则返回错误。
		仅当未读取任何字节时，错误才是EOF。如果在读取了一些但不是全部字节后发生EOF，则ReadFull返回ErrUnexpectedEOF。
		返回时，当且仅当err == nil时，n == len（buf）。
	*/
	k := make([]byte, klen)
	_, e = io.ReadFull(r, k)
	if e != nil {
		return "", nil, e
	}

	v := make([]byte, vlen)
	_, e = io.ReadFull(r, v)
	if e != nil {
		return "", nil, e
	}
	return string(k), v, nil

}

func sendResponse(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errStr := err.Error()
		tmp := fmt.Sprintf("-%d", len(errStr)) + errStr
		_, e := conn.Write([]byte(tmp))
		return e
	}
	vlen := fmt.Sprintf("%d", len(value))
	_, e := conn.Write(append([]byte(vlen), value...))
	return e
}

func (s *Server) get(conn net.Conn, r *bufio.Reader) error {
	k, e := s.readKey(r)
	if e != nil {
		return e
	}
	v,e :=s.Get(k)
	return sendResponse(v,e,conn)
}

func (s *Server) set(conn net.Conn, r *bufio.Reader) error {
	k,v, e := s.readKeyAndValue(r)
	if e != nil {
		return e
	}
	return sendResponse(nil,s.Set(k,v),conn)
}

func (s *Server) del(conn net.Conn, r *bufio.Reader) error {
	k, e := s.readKey(r)
	if e != nil {
		return e
	}
	return sendResponse(nil,s.Del(k),conn)
}
