package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer1 bytes.Buffer
	contents := "Simple byte buffer for marshaling data."
	fmt.Printf("Writing contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())

	p1 := make([]byte, 7)
	n, _ := buffer1.Read(p1)
	fmt.Printf("%d bytes read, len: %d, cap: %d\n", n, len(p1), cap(p1))
	fmt.Printf("%d bytes read, len: %d, cap: %d\n", n, buffer1.Len(), buffer1.Cap())

	contents2 := "ab"
	buffer2 := bytes.NewBufferString(contents2)
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents2, buffer2.Cap()) // 内容容器的容量为：8。
	unreadBytes := buffer2.Bytes()
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes) // 未读内容为：[97 98]。

	buffer2.WriteString("cdefg")
	fmt.Printf("The capacity of buffer %q: %d\n", buffer2, buffer2.Cap()) // 内容容器的容量仍为：8。
	unreadBytes = unreadBytes[:cap(unreadBytes)]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes) // 基于前面获取到的结果值可得，未读内容为：[97 98 99 100 101 102 103 0]。


}
