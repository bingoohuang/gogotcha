# Branchless Coding in Go

[Branchless Coding in Go](https://mattnakama.com/blog/go-branchless-coding/)

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	tests := make([]bool, 8)
	num := flag.Uint("v", 0, "flags to extract")
	flag.Parse()
	// input value, in binary
	var input uint8 = uint8(*num)
	fmt.Printf("val: %b\n", input)
	// set each boolean to a bit in the input
	if input&(1<<0) != 0 {
		tests[0] = true
	}
	if input&(1<<1) != 0 {
		tests[1] = true
	}
	if input&(1<<2) != 0 {
		tests[2] = true
	}
	if input&(1<<3) != 0 {
		tests[3] = true
	}
	if input&(1<<4) != 0 {
		tests[4] = true
	}
	if input&(1<<5) != 0 {
		tests[5] = true
	}
	if input&(1<<6) != 0 {
		tests[6] = true
	}
	if input&(1<<7) != 0 {
		tests[7] = true
	}
	fmt.Printf("result: %v\n", tests)
	for i, val := range tests {
		fmt.Printf("result %d: %t\n", i, val)
	}
}

```

`go build -o /tmp/binextract-if -gcflags='-S' cmd/branch/main.go`

```sh
🕙[2021-05-12 21:47:08.167] ❯ go build -o /tmp/binextract-if -gcflags='-S' .
# github.com/bingoohuang/gogotcha/cmd/branch
"".main STEXT size=893 args=0x0 locals=0xc0 funcid=0x0
0x0000 00000 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       TEXT    "".main(SB), ABIInternal, $192-0
0x0000 00000 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       MOVQ    (TLS), CX
0x0009 00009 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       LEAQ    -64(SP), AX
0x000e 00014 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       CMPQ    AX, 16(CX)
0x0012 00018 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       PCDATA  $0, $-2
0x0012 00018 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       JLS     883
0x0018 00024 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       PCDATA  $0, $-1
0x0018 00024 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       SUBQ    $192, SP
0x001f 00031 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       MOVQ    BP, 184(SP)
0x0027 00039 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       LEAQ    184(SP), BP
0x002f 00047 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       FUNCDATA        $0, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
0x002f 00047 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       FUNCDATA        $1, gclocals·d8210acbcf4338c24ba02b6af3d7e451(SB)
0x002f 00047 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       FUNCDATA        $2, "".main.stkobj(SB)
0x002f 00047 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:9)       LEAQ    type.bool(SB), AX
0x0036 00054 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:9)       MOVQ    AX, (SP)
0x003a 00058 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:9)       MOVQ    $8, 8(SP)
0x0043 00067 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:9)       MOVQ    $8, 16(SP)
0x004c 00076 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:9)       PCDATA  $1, $0
0x004c 00076 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:9)       CALL    runtime.makeslice(SB)
0x0051 00081 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:9)       MOVQ    24(SP), AX
0x0056 00086 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:9)       MOVQ    AX, ""..autotmp_85+112(SP)
0x005b 00091 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:10)      XCHGL   AX, AX
0x005c 00092 ($GOROOT/src/flag/flag.go:728)     MOVQ    flag.CommandLine(SB), CX
0x0063 00099 ($GOROOT/src/flag/flag.go:728)     MOVQ    CX, (SP)
0x0067 00103 ($GOROOT/src/flag/flag.go:728)     LEAQ    go.string."v"(SB), CX
0x006e 00110 ($GOROOT/src/flag/flag.go:728)     MOVQ    CX, 8(SP)
0x0073 00115 ($GOROOT/src/flag/flag.go:728)     MOVQ    $1, 16(SP)
0x007c 00124 ($GOROOT/src/flag/flag.go:728)     MOVQ    $0, 24(SP)
0x0085 00133 ($GOROOT/src/flag/flag.go:728)     LEAQ    go.string."flags to extract"(SB), CX
0x008c 00140 ($GOROOT/src/flag/flag.go:728)     MOVQ    CX, 32(SP)
0x0091 00145 ($GOROOT/src/flag/flag.go:728)     MOVQ    $16, 40(SP)
0x009a 00154 ($GOROOT/src/flag/flag.go:728)     PCDATA  $1, $1
0x009a 00154 ($GOROOT/src/flag/flag.go:728)     CALL    flag.(*FlagSet).Uint(SB)
0x009f 00159 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:11)      XCHGL   AX, AX
0x00a0 00160 ($GOROOT/src/flag/flag.go:728)     MOVQ    48(SP), AX
0x00a5 00165 ($GOROOT/src/flag/flag.go:1022)    MOVQ    os.Args+8(SB), CX
0x00ac 00172 ($GOROOT/src/flag/flag.go:1022)    MOVQ    os.Args(SB), DX
0x00b3 00179 ($GOROOT/src/flag/flag.go:1022)    MOVQ    os.Args+16(SB), BX
0x00ba 00186 ($GOROOT/src/flag/flag.go:1022)    NOP
0x00c0 00192 ($GOROOT/src/flag/flag.go:1022)    CMPQ    CX, $1
0x00c4 00196 ($GOROOT/src/flag/flag.go:1022)    JCS     872
0x00ca 00202 ($GOROOT/src/flag/flag.go:728)     MOVQ    AX, "".num+104(SP)
0x00cf 00207 ($GOROOT/src/flag/flag.go:1022)    MOVQ    flag.CommandLine(SB), AX
0x00d6 00214 ($GOROOT/src/flag/flag.go:1022)    MOVQ    AX, (SP)
0x00da 00218 ($GOROOT/src/flag/flag.go:1022)    LEAQ    -1(BX), AX
0x00de 00222 ($GOROOT/src/flag/flag.go:1022)    MOVQ    AX, BX
0x00e1 00225 ($GOROOT/src/flag/flag.go:1022)    NEGQ    AX
0x00e4 00228 ($GOROOT/src/flag/flag.go:1022)    SARQ    $63, AX
0x00e8 00232 ($GOROOT/src/flag/flag.go:1022)    ANDQ    $16, AX
0x00ec 00236 ($GOROOT/src/flag/flag.go:1022)    ADDQ    DX, AX
0x00ef 00239 ($GOROOT/src/flag/flag.go:1022)    MOVQ    AX, 8(SP)
0x00f4 00244 ($GOROOT/src/flag/flag.go:1022)    LEAQ    -1(CX), AX
0x00f8 00248 ($GOROOT/src/flag/flag.go:1022)    MOVQ    AX, 16(SP)
0x00fd 00253 ($GOROOT/src/flag/flag.go:1022)    MOVQ    BX, 24(SP)
0x0102 00258 ($GOROOT/src/flag/flag.go:1022)    PCDATA  $1, $2
0x0102 00258 ($GOROOT/src/flag/flag.go:1022)    CALL    flag.(*FlagSet).Parse(SB)
0x0107 00263 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:13)      MOVQ    "".num+104(SP), AX
0x010c 00268 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:13)      MOVQ    (AX), AX
0x010f 00271 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:13)      MOVQ    AX, ""..autotmp_86+96(SP)
0x0114 00276 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:14)      XORPS   X0, X0
0x0117 00279 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:14)      MOVUPS  X0, ""..autotmp_41+136(SP)
0x011f 00287 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:14)      LEAQ    type.uint8(SB), CX
0x0126 00294 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:14)      MOVQ    CX, ""..autotmp_41+136(SP)
0x012e 00302 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:14)      MOVBLZX AL, CX
0x0131 00305 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:14)      LEAQ    runtime.staticuint64s(SB), DX
0x0138 00312 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:14)      LEAQ    (DX)(CX*8), CX
0x013c 00316 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:14)      MOVQ    CX, ""..autotmp_41+144(SP)
0x0144 00324 (<unknown line number>)    NOP
0x0144 00324 ($GOROOT/src/fmt/print.go:213)     MOVQ    os.Stdout(SB), CX
0x014b 00331 ($GOROOT/src/fmt/print.go:213)     LEAQ    go.itab.*os.File,io.Writer(SB), BX
0x0152 00338 ($GOROOT/src/fmt/print.go:213)     MOVQ    BX, (SP)
0x0156 00342 ($GOROOT/src/fmt/print.go:213)     MOVQ    CX, 8(SP)
0x015b 00347 ($GOROOT/src/fmt/print.go:213)     LEAQ    go.string."val: %b\n"(SB), CX
0x0162 00354 ($GOROOT/src/fmt/print.go:213)     MOVQ    CX, 16(SP)
0x0167 00359 ($GOROOT/src/fmt/print.go:213)     MOVQ    $8, 24(SP)
0x0170 00368 ($GOROOT/src/fmt/print.go:213)     LEAQ    ""..autotmp_41+136(SP), CX
0x0178 00376 ($GOROOT/src/fmt/print.go:213)     MOVQ    CX, 32(SP)
0x017d 00381 ($GOROOT/src/fmt/print.go:213)     MOVQ    $1, 40(SP)
0x0186 00390 ($GOROOT/src/fmt/print.go:213)     MOVQ    $1, 48(SP)
0x018f 00399 ($GOROOT/src/fmt/print.go:213)     PCDATA  $1, $1
0x018f 00399 ($GOROOT/src/fmt/print.go:213)     CALL    fmt.Fprintf(SB)
0x0194 00404 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:16)      MOVQ    ""..autotmp_86+96(SP), AX
0x0199 00409 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:16)      NOP
0x01a0 00416 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:16)      TESTB   $1, AL
0x01a2 00418 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:16)      JEQ     862
0x01a8 00424 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:17)      MOVQ    ""..autotmp_85+112(SP), CX
0x01ad 00429 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:17)      MOVB    $1, (CX)
0x01b0 00432 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:19)      TESTB   $2, AL
0x01b2 00434 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:19)      JEQ     440
0x01b4 00436 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:20)      MOVB    $1, 1(CX)
0x01b8 00440 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:22)      TESTB   $4, AL
0x01ba 00442 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:22)      JEQ     448
0x01bc 00444 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:23)      MOVB    $1, 2(CX)
0x01c0 00448 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:25)      TESTB   $8, AL
0x01c2 00450 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:25)      JEQ     456
0x01c4 00452 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:26)      MOVB    $1, 3(CX)
0x01c8 00456 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:28)      TESTB   $16, AL
0x01ca 00458 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:28)      JEQ     464
0x01cc 00460 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:29)      MOVB    $1, 4(CX)
0x01d0 00464 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:31)      TESTB   $32, AL
0x01d2 00466 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:31)      JEQ     472
0x01d4 00468 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:32)      MOVB    $1, 5(CX)
0x01d8 00472 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:34)      TESTB   $64, AL
0x01da 00474 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:34)      JEQ     480
0x01dc 00476 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:35)      MOVB    $1, 6(CX)
0x01e0 00480 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:37)      TESTB   $-128, AL
0x01e2 00482 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:37)      JEQ     488
0x01e4 00484 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:38)      MOVB    $1, 7(CX)
0x01e8 00488 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:40)      MOVQ    CX, (SP)
0x01ec 00492 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:40)      MOVQ    $8, 8(SP)
0x01f5 00501 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:40)      MOVQ    $8, 16(SP)
0x01fe 00510 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:40)      NOP
0x0200 00512 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:40)      CALL    runtime.convTslice(SB)
0x0205 00517 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:40)      MOVQ    24(SP), AX
0x020a 00522 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:40)      XORPS   X0, X0
0x020d 00525 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:40)      MOVUPS  X0, ""..autotmp_46+120(SP)
0x0212 00530 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:40)      LEAQ    type.[]bool(SB), CX
0x0219 00537 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:40)      MOVQ    CX, ""..autotmp_46+120(SP)
0x021e 00542 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:40)      MOVQ    AX, ""..autotmp_46+128(SP)
0x0226 00550 (<unknown line number>)    NOP
0x0226 00550 ($GOROOT/src/fmt/print.go:213)     MOVQ    os.Stdout(SB), AX
0x022d 00557 ($GOROOT/src/fmt/print.go:213)     LEAQ    go.itab.*os.File,io.Writer(SB), CX
0x0234 00564 ($GOROOT/src/fmt/print.go:213)     MOVQ    CX, (SP)
0x0238 00568 ($GOROOT/src/fmt/print.go:213)     MOVQ    AX, 8(SP)
0x023d 00573 ($GOROOT/src/fmt/print.go:213)     LEAQ    go.string."result: %v\n"(SB), AX
0x0244 00580 ($GOROOT/src/fmt/print.go:213)     MOVQ    AX, 16(SP)
0x0249 00585 ($GOROOT/src/fmt/print.go:213)     MOVQ    $11, 24(SP)
0x0252 00594 ($GOROOT/src/fmt/print.go:213)     LEAQ    ""..autotmp_46+120(SP), AX
0x0257 00599 ($GOROOT/src/fmt/print.go:213)     MOVQ    AX, 32(SP)
0x025c 00604 ($GOROOT/src/fmt/print.go:213)     MOVQ    $1, 40(SP)
0x0265 00613 ($GOROOT/src/fmt/print.go:213)     MOVQ    $1, 48(SP)
0x026e 00622 ($GOROOT/src/fmt/print.go:213)     CALL    fmt.Fprintf(SB)
0x0273 00627 ($GOROOT/src/fmt/print.go:213)     XORL    AX, AX
0x0275 00629 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      JMP     836
0x027a 00634 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      MOVQ    AX, "".i+80(SP)
0x027f 00639 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      MOVQ    ""..autotmp_85+112(SP), CX
0x0284 00644 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      MOVBLZX (CX)(AX*1), DX
0x0288 00648 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      MOVQ    DX, ""..autotmp_87+88(SP)
0x028d 00653 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      MOVQ    AX, (SP)
0x0291 00657 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      CALL    runtime.convT64(SB)
0x0296 00662 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      MOVQ    8(SP), AX
0x029b 00667 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      XORPS   X0, X0
0x029e 00670 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      MOVUPS  X0, ""..autotmp_54+152(SP)
0x02a6 00678 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      MOVUPS  X0, ""..autotmp_54+168(SP)
0x02ae 00686 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      LEAQ    type.int(SB), CX
0x02b5 00693 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      MOVQ    CX, ""..autotmp_54+152(SP)
0x02bd 00701 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      MOVQ    AX, ""..autotmp_54+160(SP)
0x02c5 00709 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      LEAQ    type.bool(SB), AX
0x02cc 00716 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      MOVQ    AX, ""..autotmp_54+168(SP)
0x02d4 00724 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      MOVQ    ""..autotmp_87+88(SP), DX
0x02d9 00729 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      LEAQ    runtime.staticuint64s(SB), BX
0x02e0 00736 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      LEAQ    (BX)(DX*8), DX
0x02e4 00740 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:42)      MOVQ    DX, ""..autotmp_54+176(SP)
0x02ec 00748 (<unknown line number>)    NOP
0x02ec 00748 ($GOROOT/src/fmt/print.go:213)     MOVQ    os.Stdout(SB), DX
0x02f3 00755 ($GOROOT/src/fmt/print.go:213)     LEAQ    go.itab.*os.File,io.Writer(SB), SI
0x02fa 00762 ($GOROOT/src/fmt/print.go:213)     MOVQ    SI, (SP)
0x02fe 00766 ($GOROOT/src/fmt/print.go:213)     MOVQ    DX, 8(SP)
0x0303 00771 ($GOROOT/src/fmt/print.go:213)     LEAQ    go.string."result %d: %t\n"(SB), DX
0x030a 00778 ($GOROOT/src/fmt/print.go:213)     MOVQ    DX, 16(SP)
0x030f 00783 ($GOROOT/src/fmt/print.go:213)     MOVQ    $14, 24(SP)
0x0318 00792 ($GOROOT/src/fmt/print.go:213)     LEAQ    ""..autotmp_54+152(SP), DI
0x0320 00800 ($GOROOT/src/fmt/print.go:213)     MOVQ    DI, 32(SP)
0x0325 00805 ($GOROOT/src/fmt/print.go:213)     MOVQ    $2, 40(SP)
0x032e 00814 ($GOROOT/src/fmt/print.go:213)     MOVQ    $2, 48(SP)
0x0337 00823 ($GOROOT/src/fmt/print.go:213)     CALL    fmt.Fprintf(SB)
0x033c 00828 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      MOVQ    "".i+80(SP), AX
0x0341 00833 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      INCQ    AX
0x0344 00836 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      CMPQ    AX, $8
0x0348 00840 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      JLT     634
0x034e 00846 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      PCDATA  $1, $-1
0x034e 00846 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      MOVQ    184(SP), BP
0x0356 00854 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      ADDQ    $192, SP
0x035d 00861 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:41)      RET
0x035e 00862 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:20)      MOVQ    ""..autotmp_85+112(SP), CX
0x0363 00867 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:16)      JMP     432
0x0368 00872 ($GOROOT/src/flag/flag.go:1022)    MOVL    $1, AX
0x036d 00877 ($GOROOT/src/flag/flag.go:1022)    PCDATA  $1, $0
0x036d 00877 ($GOROOT/src/flag/flag.go:1022)    CALL    runtime.panicSliceB(SB)
0x0372 00882 ($GOROOT/src/flag/flag.go:1022)    XCHGL   AX, AX
0x0373 00883 ($GOROOT/src/flag/flag.go:1022)    NOP
0x0373 00883 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       PCDATA  $1, $-1
0x0373 00883 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       PCDATA  $0, $-2
0x0373 00883 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       CALL    runtime.morestack_noctxt(SB)
0x0378 00888 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       PCDATA  $0, $-1
0x0378 00888 (/Users/bingoo/GitHub/gogotcha/cmd/branch/main.go:8)       JMP     0
0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 44 24 c0 48 3b  eH..%....H.D$.H;
0x0010 41 10 0f 86 5b 03 00 00 48 81 ec c0 00 00 00 48  A...[...H......H
0x0020 89 ac 24 b8 00 00 00 48 8d ac 24 b8 00 00 00 48  ..$....H..$....H
0x0030 8d 05 00 00 00 00 48 89 04 24 48 c7 44 24 08 08  ......H..$H.D$..
0x0040 00 00 00 48 c7 44 24 10 08 00 00 00 e8 00 00 00  ...H.D$.........
0x0050 00 48 8b 44 24 18 48 89 44 24 70 90 48 8b 0d 00  .H.D$.H.D$p.H...
0x0060 00 00 00 48 89 0c 24 48 8d 0d 00 00 00 00 48 89  ...H..$H......H.
0x0070 4c 24 08 48 c7 44 24 10 01 00 00 00 48 c7 44 24  L$.H.D$.....H.D$
0x0080 18 00 00 00 00 48 8d 0d 00 00 00 00 48 89 4c 24  .....H......H.L$
0x0090 20 48 c7 44 24 28 10 00 00 00 e8 00 00 00 00 90   H.D$(..........
0x00a0 48 8b 44 24 30 48 8b 0d 00 00 00 00 48 8b 15 00  H.D$0H......H...
0x00b0 00 00 00 48 8b 1d 00 00 00 00 66 0f 1f 44 00 00  ...H......f..D..
0x00c0 48 83 f9 01 0f 82 9e 02 00 00 48 89 44 24 68 48  H.........H.D$hH
0x00d0 8b 05 00 00 00 00 48 89 04 24 48 8d 43 ff 48 89  ......H..$H.C.H.
0x00e0 c3 48 f7 d8 48 c1 f8 3f 48 83 e0 10 48 01 d0 48  .H..H..?H...H..H
0x00f0 89 44 24 08 48 8d 41 ff 48 89 44 24 10 48 89 5c  .D$.H.A.H.D$.H.\
0x0100 24 18 e8 00 00 00 00 48 8b 44 24 68 48 8b 00 48  $......H.D$hH..H
0x0110 89 44 24 60 0f 57 c0 0f 11 84 24 88 00 00 00 48  .D$`.W....$....H
0x0120 8d 0d 00 00 00 00 48 89 8c 24 88 00 00 00 0f b6  ......H..$......
0x0130 c8 48 8d 15 00 00 00 00 48 8d 0c ca 48 89 8c 24  .H......H...H..$
0x0140 90 00 00 00 48 8b 0d 00 00 00 00 48 8d 1d 00 00  ....H......H....
0x0150 00 00 48 89 1c 24 48 89 4c 24 08 48 8d 0d 00 00  ..H..$H.L$.H....
0x0160 00 00 48 89 4c 24 10 48 c7 44 24 18 08 00 00 00  ..H.L$.H.D$.....
0x0170 48 8d 8c 24 88 00 00 00 48 89 4c 24 20 48 c7 44  H..$....H.L$ H.D
0x0180 24 28 01 00 00 00 48 c7 44 24 30 01 00 00 00 e8  $(....H.D$0.....
0x0190 00 00 00 00 48 8b 44 24 60 0f 1f 80 00 00 00 00  ....H.D$`.......
0x01a0 a8 01 0f 84 b6 01 00 00 48 8b 4c 24 70 c6 01 01  ........H.L$p...
0x01b0 a8 02 74 04 c6 41 01 01 a8 04 74 04 c6 41 02 01  ..t..A....t..A..
0x01c0 a8 08 74 04 c6 41 03 01 a8 10 74 04 c6 41 04 01  ..t..A....t..A..
0x01d0 a8 20 74 04 c6 41 05 01 a8 40 74 04 c6 41 06 01  . t..A...@t..A..
0x01e0 a8 80 74 04 c6 41 07 01 48 89 0c 24 48 c7 44 24  ..t..A..H..$H.D$
0x01f0 08 08 00 00 00 48 c7 44 24 10 08 00 00 00 66 90  .....H.D$.....f.
0x0200 e8 00 00 00 00 48 8b 44 24 18 0f 57 c0 0f 11 44  .....H.D$..W...D
0x0210 24 78 48 8d 0d 00 00 00 00 48 89 4c 24 78 48 89  $xH......H.L$xH.
0x0220 84 24 80 00 00 00 48 8b 05 00 00 00 00 48 8d 0d  .$....H......H..
0x0230 00 00 00 00 48 89 0c 24 48 89 44 24 08 48 8d 05  ....H..$H.D$.H..
0x0240 00 00 00 00 48 89 44 24 10 48 c7 44 24 18 0b 00  ....H.D$.H.D$...
0x0250 00 00 48 8d 44 24 78 48 89 44 24 20 48 c7 44 24  ..H.D$xH.D$ H.D$
0x0260 28 01 00 00 00 48 c7 44 24 30 01 00 00 00 e8 00  (....H.D$0......
0x0270 00 00 00 31 c0 e9 ca 00 00 00 48 89 44 24 50 48  ...1......H.D$PH
0x0280 8b 4c 24 70 0f b6 14 01 48 89 54 24 58 48 89 04  .L$p....H.T$XH..
0x0290 24 e8 00 00 00 00 48 8b 44 24 08 0f 57 c0 0f 11  $.....H.D$..W...
0x02a0 84 24 98 00 00 00 0f 11 84 24 a8 00 00 00 48 8d  .$.......$....H.
0x02b0 0d 00 00 00 00 48 89 8c 24 98 00 00 00 48 89 84  .....H..$....H..
0x02c0 24 a0 00 00 00 48 8d 05 00 00 00 00 48 89 84 24  $....H......H..$
0x02d0 a8 00 00 00 48 8b 54 24 58 48 8d 1d 00 00 00 00  ....H.T$XH......
0x02e0 48 8d 14 d3 48 89 94 24 b0 00 00 00 48 8b 15 00  H...H..$....H...
0x02f0 00 00 00 48 8d 35 00 00 00 00 48 89 34 24 48 89  ...H.5....H.4$H.
0x0300 54 24 08 48 8d 15 00 00 00 00 48 89 54 24 10 48  T$.H......H.T$.H
0x0310 c7 44 24 18 0e 00 00 00 48 8d bc 24 98 00 00 00  .D$.....H..$....
0x0320 48 89 7c 24 20 48 c7 44 24 28 02 00 00 00 48 c7  H.|$ H.D$(....H.
0x0330 44 24 30 02 00 00 00 e8 00 00 00 00 48 8b 44 24  D$0.........H.D$
0x0340 50 48 ff c0 48 83 f8 08 0f 8c 2c ff ff ff 48 8b  PH..H.....,...H.
0x0350 ac 24 b8 00 00 00 48 81 c4 c0 00 00 00 c3 48 8b  .$....H.......H.
0x0360 4c 24 70 e9 48 fe ff ff b8 01 00 00 00 e8 00 00  L$p.H...........
0x0370 00 00 90 e8 00 00 00 00 e9 83 fc ff ff           .............
rel 3+0 t=25 type.uint8+0
rel 3+0 t=25 type.*os.File+0
rel 3+0 t=25 type.[]bool+0
rel 3+0 t=25 type.*os.File+0
rel 3+0 t=25 type.int+0
rel 3+0 t=25 type.bool+0
rel 3+0 t=25 type.*os.File+0
rel 5+4 t=17 TLS+0
rel 50+4 t=16 type.bool+0
rel 77+4 t=8 runtime.makeslice+0
rel 95+4 t=16 flag.CommandLine+0
rel 106+4 t=16 go.string."v"+0
rel 136+4 t=16 go.string."flags to extract"+0
rel 155+4 t=8 flag.(*FlagSet).Uint+0
rel 168+4 t=16 os.Args+8
rel 175+4 t=16 os.Args+0
rel 182+4 t=16 os.Args+16
rel 210+4 t=16 flag.CommandLine+0
rel 259+4 t=8 flag.(*FlagSet).Parse+0
rel 290+4 t=16 type.uint8+0
rel 308+4 t=16 runtime.staticuint64s+0
rel 327+4 t=16 os.Stdout+0
rel 334+4 t=16 go.itab.*os.File,io.Writer+0
rel 350+4 t=16 go.string."val: %b\n"+0
rel 400+4 t=8 fmt.Fprintf+0
rel 513+4 t=8 runtime.convTslice+0
rel 533+4 t=16 type.[]bool+0
rel 553+4 t=16 os.Stdout+0
rel 560+4 t=16 go.itab.*os.File,io.Writer+0
rel 576+4 t=16 go.string."result: %v\n"+0
rel 623+4 t=8 fmt.Fprintf+0
rel 658+4 t=8 runtime.convT64+0
rel 689+4 t=16 type.int+0
rel 712+4 t=16 type.bool+0
rel 732+4 t=16 runtime.staticuint64s+0
rel 751+4 t=16 os.Stdout+0
rel 758+4 t=16 go.itab.*os.File,io.Writer+0
rel 774+4 t=16 go.string."result %d: %t\n"+0
rel 824+4 t=8 fmt.Fprintf+0
rel 878+4 t=8 runtime.panicSliceB+0
rel 884+4 t=8 runtime.morestack_noctxt+0
os.(*File).close STEXT dupok nosplit size=26 args=0x18 locals=0x0 funcid=0x0
0x0000 00000 (<autogenerated>:1)        TEXT    os.(*File).close(SB), DUPOK|NOSPLIT|ABIInternal, $0-24
0x0000 00000 (<autogenerated>:1)        FUNCDATA        $0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
0x0000 00000 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
0x0000 00000 (<autogenerated>:1)        MOVQ    ""..this+8(SP), AX
0x0005 00005 (<autogenerated>:1)        MOVQ    (AX), AX
0x0008 00008 (<autogenerated>:1)        MOVQ    AX, ""..this+8(SP)
0x000d 00013 (<autogenerated>:1)        XORPS   X0, X0
0x0010 00016 (<autogenerated>:1)        MOVUPS  X0, "".~r0+16(SP)
0x0015 00021 (<autogenerated>:1)        JMP     os.(*file).close(SB)
0x0000 48 8b 44 24 08 48 8b 00 48 89 44 24 08 0f 57 c0  H.D$.H..H.D$..W.
0x0010 0f 11 44 24 10 e9 00 00 00 00                    ..D$......
rel 22+4 t=8 os.(*file).close+0
type..eq.[2]interface {} STEXT dupok size=170 args=0x18 locals=0x30 funcid=0x0
0x0000 00000 (<autogenerated>:1)        TEXT    type..eq.[2]interface {}(SB), DUPOK|ABIInternal, $48-24
0x0000 00000 (<autogenerated>:1)        MOVQ    (TLS), CX
0x0009 00009 (<autogenerated>:1)        CMPQ    SP, 16(CX)
0x000d 00013 (<autogenerated>:1)        PCDATA  $0, $-2
0x000d 00013 (<autogenerated>:1)        JLS     160
0x0013 00019 (<autogenerated>:1)        PCDATA  $0, $-1
0x0013 00019 (<autogenerated>:1)        SUBQ    $48, SP
0x0017 00023 (<autogenerated>:1)        MOVQ    BP, 40(SP)
0x001c 00028 (<autogenerated>:1)        LEAQ    40(SP), BP
0x0021 00033 (<autogenerated>:1)        FUNCDATA        $0, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
0x0021 00033 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
0x0021 00033 (<autogenerated>:1)        MOVQ    "".q+64(SP), AX
0x0026 00038 (<autogenerated>:1)        MOVQ    "".p+56(SP), CX
0x002b 00043 (<autogenerated>:1)        XORL    DX, DX
0x002d 00045 (<autogenerated>:1)        JMP     66
0x002f 00047 (<autogenerated>:1)        MOVQ    ""..autotmp_6+32(SP), BX
0x0034 00052 (<autogenerated>:1)        LEAQ    1(BX), DX
0x0038 00056 (<autogenerated>:1)        MOVQ    "".q+64(SP), AX
0x003d 00061 (<autogenerated>:1)        MOVQ    "".p+56(SP), CX
0x0042 00066 (<autogenerated>:1)        CMPQ    DX, $2
0x0046 00070 (<autogenerated>:1)        JGE     149
0x0048 00072 (<autogenerated>:1)        MOVQ    DX, BX
0x004b 00075 (<autogenerated>:1)        SHLQ    $4, DX
0x004f 00079 (<autogenerated>:1)        MOVQ    (CX)(DX*1), SI
0x0053 00083 (<autogenerated>:1)        MOVQ    (AX)(DX*1), DI
0x0057 00087 (<autogenerated>:1)        MOVQ    8(DX)(CX*1), R8
0x005c 00092 (<autogenerated>:1)        MOVQ    8(DX)(AX*1), DX
0x0061 00097 (<autogenerated>:1)        CMPQ    DI, SI
0x0064 00100 (<autogenerated>:1)        JNE     133
0x0066 00102 (<autogenerated>:1)        MOVQ    BX, ""..autotmp_6+32(SP)
0x006b 00107 (<autogenerated>:1)        MOVQ    SI, (SP)
0x006f 00111 (<autogenerated>:1)        MOVQ    R8, 8(SP)
0x0074 00116 (<autogenerated>:1)        MOVQ    DX, 16(SP)
0x0079 00121 (<autogenerated>:1)        PCDATA  $1, $0
0x0079 00121 (<autogenerated>:1)        CALL    runtime.efaceeq(SB)
0x007e 00126 (<autogenerated>:1)        CMPB    24(SP), $0
0x0083 00131 (<autogenerated>:1)        JNE     47
0x0085 00133 (<autogenerated>:1)        XORL    AX, AX
0x0087 00135 (<autogenerated>:1)        MOVB    AL, "".r+72(SP)
0x008b 00139 (<autogenerated>:1)        MOVQ    40(SP), BP
0x0090 00144 (<autogenerated>:1)        ADDQ    $48, SP
0x0094 00148 (<autogenerated>:1)        RET
0x0095 00149 (<autogenerated>:1)        MOVL    $1, AX
0x009a 00154 (<autogenerated>:1)        JMP     135
0x009c 00156 (<autogenerated>:1)        NOP
0x009c 00156 (<autogenerated>:1)        PCDATA  $1, $-1
0x009c 00156 (<autogenerated>:1)        PCDATA  $0, $-2
0x009c 00156 (<autogenerated>:1)        NOP
0x00a0 00160 (<autogenerated>:1)        CALL    runtime.morestack_noctxt(SB)
0x00a5 00165 (<autogenerated>:1)        PCDATA  $0, $-1
0x00a5 00165 (<autogenerated>:1)        JMP     0
0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 8d  eH..%....H;a....
0x0010 00 00 00 48 83 ec 30 48 89 6c 24 28 48 8d 6c 24  ...H..0H.l$(H.l$
0x0020 28 48 8b 44 24 40 48 8b 4c 24 38 31 d2 eb 13 48  (H.D$@H.L$81...H
0x0030 8b 5c 24 20 48 8d 53 01 48 8b 44 24 40 48 8b 4c  .\$ H.S.H.D$@H.L
0x0040 24 38 48 83 fa 02 7d 4d 48 89 d3 48 c1 e2 04 48  $8H...}MH..H...H
0x0050 8b 34 11 48 8b 3c 10 4c 8b 44 0a 08 48 8b 54 02  .4.H.<.L.D..H.T.
0x0060 08 48 39 f7 75 1f 48 89 5c 24 20 48 89 34 24 4c  .H9.u.H.\$ H.4$L
0x0070 89 44 24 08 48 89 54 24 10 e8 00 00 00 00 80 7c  .D$.H.T$.......|
0x0080 24 18 00 75 aa 31 c0 88 44 24 48 48 8b 6c 24 28  $..u.1..D$HH.l$(
0x0090 48 83 c4 30 c3 b8 01 00 00 00 eb eb 0f 1f 40 00  H..0..........@.
0x00a0 e8 00 00 00 00 e9 56 ff ff ff                    ......V...
rel 5+4 t=17 TLS+0
rel 122+4 t=8 runtime.efaceeq+0
rel 161+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename.main SDWARFCUINFO dupok size=0
0x0000 6d 61 69 6e                                      main
go.string..gostring.125.1800acee410835f7cb749378ba445f9dbb9c2a2209c6b3810070b8f798227d1c SRODATA dupok size=125
0x0000 30 77 af 0c 92 74 08 02 41 e1 c1 07 e6 d6 18 e6  0w...t..A.......
0x0010 70 61 74 68 09 67 69 74 68 75 62 2e 63 6f 6d 2f  path.github.com/
0x0020 62 69 6e 67 6f 6f 68 75 61 6e 67 2f 67 6f 67 6f  bingoohuang/gogo
0x0030 74 63 68 61 2f 63 6d 64 2f 62 72 61 6e 63 68 0a  tcha/cmd/branch.
0x0040 6d 6f 64 09 67 69 74 68 75 62 2e 63 6f 6d 2f 62  mod.github.com/b
0x0050 69 6e 67 6f 6f 68 75 61 6e 67 2f 67 6f 67 6f 74  ingoohuang/gogot
0x0060 63 68 61 09 28 64 65 76 65 6c 29 09 0a f9 32 43  cha.(devel)...2C
0x0070 31 86 18 20 72 00 82 42 10 41 16 d8 f2           1.. r..B.A...
""..inittask SNOPTRDATA size=40
0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0020 00 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 flag..inittask+0
rel 32+8 t=1 fmt..inittask+0
go.info.flag.Uint$abstract SDWARFABSFCN dupok size=49
0x0000 04 66 6c 61 67 2e 55 69 6e 74 00 01 01 11 6e 61  .flag.Uint....na
0x0010 6d 65 00 00 00 00 00 00 11 76 61 6c 75 65 00 00  me.......value..
0x0020 00 00 00 00 11 75 73 61 67 65 00 00 00 00 00 00  .....usage......
0x0030 00                                               .
rel 0+0 t=24 type.string+0
rel 0+0 t=24 type.uint+0
rel 20+4 t=31 go.info.string+0
rel 32+4 t=31 go.info.uint+0
rel 44+4 t=31 go.info.string+0
go.info.flag.Parse$abstract SDWARFABSFCN dupok size=15
0x0000 04 66 6c 61 67 2e 50 61 72 73 65 00 01 01 00     .flag.Parse....
go.info.fmt.Printf$abstract SDWARFABSFCN dupok size=54
0x0000 04 66 6d 74 2e 50 72 69 6e 74 66 00 01 01 11 66  .fmt.Printf....f
0x0010 6f 72 6d 61 74 00 00 00 00 00 00 11 61 00 00 00  ormat.......a...
0x0020 00 00 00 11 6e 00 01 00 00 00 00 11 65 72 72 00  ....n.......err.
0x0030 01 00 00 00 00 00                                ......
rel 0+0 t=24 type.[]interface {}+0
rel 0+0 t=24 type.error+0
rel 0+0 t=24 type.int+0
rel 0+0 t=24 type.string+0
rel 23+4 t=31 go.info.string+0
rel 31+4 t=31 go.info.[]interface {}+0
rel 39+4 t=31 go.info.int+0
rel 49+4 t=31 go.info.error+0
go.string."v" SRODATA dupok size=1
0x0000 76                                               v
go.string."flags to extract" SRODATA dupok size=16
0x0000 66 6c 61 67 73 20 74 6f 20 65 78 74 72 61 63 74  flags to extract
go.string."val: %b\n" SRODATA dupok size=8
0x0000 76 61 6c 3a 20 25 62 0a                          val: %b.
go.string."result: %v\n" SRODATA dupok size=11
0x0000 72 65 73 75 6c 74 3a 20 25 76 0a                 result: %v.
go.string."result %d: %t\n" SRODATA dupok size=14
0x0000 72 65 73 75 6c 74 20 25 64 3a 20 25 74 0a        result %d: %t.
runtime.nilinterequal·f SRODATA dupok size=8
0x0000 00 00 00 00 00 00 00 00                          ........
rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
0x0000 00 00 00 00 00 00 00 00                          ........
rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 runtime.memequal64·f+0
rel 32+8 t=1 runtime.gcbits.01+0
rel 40+4 t=5 type..namedata.*interface {}-+0
rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
0x0000 02                                               .
type.interface {} SRODATA dupok size=80
0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
rel 24+8 t=1 runtime.nilinterequal·f+0
rel 32+8 t=1 runtime.gcbits.02+0
rel 40+4 t=5 type..namedata.*interface {}-+0
rel 44+4 t=6 type.*interface {}+0
rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface
0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 runtime.memequal64·f+0
rel 32+8 t=1 runtime.gcbits.01+0
rel 40+4 t=5 type..namedata.*[]interface {}-+0
rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 32+8 t=1 runtime.gcbits.01+0
rel 40+4 t=5 type..namedata.*[]interface {}-+0
rel 44+4 t=6 type.*[]interface {}+0
rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 runtime.memequal64·f+0
rel 32+8 t=1 runtime.gcbits.01+0
rel 40+4 t=5 type..namedata.*[1]interface {}-+0
rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0040 01 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 runtime.nilinterequal·f+0
rel 32+8 t=1 runtime.gcbits.02+0
rel 40+4 t=5 type..namedata.*[1]interface {}-+0
rel 44+4 t=6 type.*[1]interface {}+0
rel 48+8 t=1 type.interface {}+0
rel 56+8 t=1 type.[]interface {}+0
type..eqfunc.[2]interface {} SRODATA dupok size=8
0x0000 00 00 00 00 00 00 00 00                          ........
rel 0+8 t=1 type..eq.[2]interface {}+0
type..namedata.*[2]interface {}- SRODATA dupok size=19
0x0000 00 00 10 2a 5b 32 5d 69 6e 74 65 72 66 61 63 65  ...*[2]interface
0x0010 20 7b 7d                                          {}
type.*[2]interface {} SRODATA dupok size=56
0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
0x0010 be 73 2d 71 08 08 08 36 00 00 00 00 00 00 00 00  .s-q...6........
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 runtime.memequal64·f+0
rel 32+8 t=1 runtime.gcbits.01+0
rel 40+4 t=5 type..namedata.*[2]interface {}-+0
rel 48+8 t=1 type.[2]interface {}+0
runtime.gcbits.0a SRODATA dupok size=1
0x0000 0a                                               .
type.[2]interface {} SRODATA dupok size=72
0x0000 20 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00   ....... .......
0x0010 2c 59 a4 f1 02 08 08 11 00 00 00 00 00 00 00 00  ,Y..............
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0040 02 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 type..eqfunc.[2]interface {}+0
rel 32+8 t=1 runtime.gcbits.0a+0
rel 40+4 t=5 type..namedata.*[2]interface {}-+0
rel 44+4 t=6 type.*[2]interface {}+0
rel 48+8 t=1 type.interface {}+0
rel 56+8 t=1 type.[]interface {}+0
runtime.modinfo SDATA size=16
0x0000 00 00 00 00 00 00 00 00 7d 00 00 00 00 00 00 00  ........}.......
rel 0+8 t=1 go.string..gostring.125.1800acee410835f7cb749378ba445f9dbb9c2a2209c6b3810070b8f798227d1c+0
type..namedata.*[]bool- SRODATA dupok size=10
0x0000 00 00 07 2a 5b 5d 62 6f 6f 6c                    ...*[]bool
type.*[]bool SRODATA dupok size=56
0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
0x0010 57 a9 a7 bd 08 08 08 36 00 00 00 00 00 00 00 00  W......6........
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 runtime.memequal64·f+0
rel 32+8 t=1 runtime.gcbits.01+0
rel 40+4 t=5 type..namedata.*[]bool-+0
rel 48+8 t=1 type.[]bool+0
type.[]bool SRODATA dupok size=56
0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
0x0010 b1 e5 81 e7 02 08 08 17 00 00 00 00 00 00 00 00  ................
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 32+8 t=1 runtime.gcbits.01+0
rel 40+4 t=5 type..namedata.*[]bool-+0
rel 44+4 t=6 type.*[]bool+0
rel 48+8 t=1 type.bool+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
rel 0+8 t=1 type.io.Writer+0
rel 8+8 t=1 type.*os.File+0
rel 24+8 t=1 os.(*File).Write+0
type..importpath.flag. SRODATA dupok size=7
0x0000 00 00 04 66 6c 61 67                             ...flag
type..importpath.fmt. SRODATA dupok size=6
0x0000 00 00 03 66 6d 74                                ...fmt
type..importpath.unsafe. SRODATA dupok size=9
0x0000 00 00 06 75 6e 73 61 66 65                       ...unsafe
gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·d8210acbcf4338c24ba02b6af3d7e451 SRODATA dupok size=14
0x0000 03 00 00 00 0a 00 00 00 00 00 02 00 03 00        ..............
"".main.stkobj SRODATA static size=56
0x0000 03 00 00 00 00 00 00 00 c0 ff ff ff ff ff ff ff  ................
0x0010 00 00 00 00 00 00 00 00 d0 ff ff ff ff ff ff ff  ................
0x0020 00 00 00 00 00 00 00 00 e0 ff ff ff ff ff ff ff  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 16+8 t=1 type.[1]interface {}+0
rel 32+8 t=1 type.[1]interface {}+0
rel 48+8 t=1 type.[2]interface {}+0
gclocals·e6397a44f8e1b6e77d0f200b4fba5269 SRODATA dupok size=10
0x0000 02 00 00 00 03 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·dc9b0298814590ca3ffc3a889546fc8b SRODATA dupok size=10
0x0000 02 00 00 00 02 00 00 00 03 00                    ..........
```


| -  | meaning              |
|------|------------------------------------------------------|
| MOV  | move value on left to location on right              |
| NOP  | no operation (do nothing)                            |
| TEST | do a boolean AND comparison between left and right   |
| JEQ  | if last operation resulted in zero, jump to location |
| $    | constant. $1==0b0001, $2==0b0010, $4==0b0100, etc.   |
| AX   | cpu register; used for comparisons here              |
| AL   | lower 8 bits of AX                                   |
| CX   | cpu register; stores address of the output array     |
