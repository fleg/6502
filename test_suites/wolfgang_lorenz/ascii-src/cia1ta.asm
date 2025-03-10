
           *= $0801
           .byte $4c,$14,$08,$00,$97
turboass   = 780
           .text "780"
           .byte $2c,$30,$3a,$9e,$32,$30
           .byte $37,$33,$00,$00,$00
           lda #1
           sta turboass
           jmp main


print
           .block
           pla
           sta print0+1
           pla
           sta print0+2
           ldx #1
print0
           lda !*,x
           beq print1
           jsr $ffd2
           inx
           bne print0
print1
           sec
           txa
           adc print0+1
           sta print2+1
           lda #0
           adc print0+2
           sta print2+2
print2
           jmp !*
           .bend


printhb
           .block
           pha
           lsr a
           lsr a
           lsr a
           lsr a
           jsr printhn
           pla
           and #$0f
printhn
           ora #$30
           cmp #$3a
           bcc printhn0
           adc #6
printhn0
           jsr $ffd2
           rts
           .bend


waitborder
           .block
           lda $d011
           bmi ok
wait
           lda $d012
           cmp #30
           bcs wait
ok
           rts
           .bend


waitkey
           .block
           jsr $fda3
wait
           jsr $ffe4
           beq wait
           cmp #3
           beq stop
           rts
stop
           lda turboass
           beq basic
           jmp $8000
basic
           jmp $a474
           .bend


report
           .block
           sta savea+1
           stx savex+1
           sty savey+1
           lda #13
           jsr $ffd2
           jsr $ffd2
           pla
           sta print0+1
           pla
           sta print0+2
           ldx #1
print0
           lda !*,x
           beq print1
           jsr $ffd2
           inx
           bne print0
print1
           sec
           txa
           adc print0+1
           sta print2+1
           lda #0
           adc print0+2
           sta print2+2
           jsr print
           .byte 13
           .text "idx"
           .byte 0
savex
           lda #$11
           jsr printhb
           lda #32
           jsr $ffd2
savea
           lda #$11
           jsr printhb
           jsr print
           .byte 13
           .text "right "
           .byte 0
savey
           lda #$11
           jsr printhb
           jsr waitkey
print2
           jmp !*
           .bend



;combinations tested
;before
;  4/6 0 latch low
;  4/6 1 latch low
;  4/6 2 latch low
;  4/6 7 latch low
;  e/f 0 stopped/running
;  e/f 3 continuous/one shot
;after
;  4/6 0 latch low
;  4/6 1 latch low
;  4/6 2 latch low
;  e/f 0 stopped/running
;  e/f 3 continuous/one shot
;  e/f 4 no force load/force load
;check
;  4/6 counter low
;  5/7 counter high
;  d/d icr
;  e/f control

ieindex    .byte 0
beindex    .byte 0
betab      .byte $00,$01,$08,$09
ietab      .byte $10,$11,$18,$19
i4         .byte 0
ie         .byte 0
b4         .byte 0
be         .byte 0
a4         .byte 0
ad         .byte 0
ae         .byte 0
r4         .byte 0
rd         .byte 0
re         .byte 0

main
           jsr print
           .byte 13
           .text "(up)cia1ta"
           .byte 0

           .block
           lda #30
           sta i4
           lda #20
           sta b4
           lda #0
           sta ieindex
           sta beindex
loop
           sei
           lda #$7f
           sta $dc0d
           lda #$81
           sta $dc0d
           jsr waitborder
           lda #0
           sta $dc0e
           sta $dc05
           ldx ieindex
           lda ietab,x
           sta ie
           ldx beindex
           ldy betab,x
           sty be
           ldx i4
           stx $dc04
           ldx #$10
           stx $dc0e
           bit $dc0d
           sta $dc0e
           lda b4
           sta $dc04
           sty $dc0e
           lda $dc04
           ldx $dc0d
           ldy $dc0e
           sta a4
           stx ad
           sty ae
           jsr $fda3
           cli

           lda ieindex
           asl a
           asl a
           asl a
           ora beindex
           asl a
           tax
           lda jumptab+0,x
           sta 172
           lda jumptab+1,x
           sta 173
           jmp (172)

jumptab    .word x000,x001,x008,x009
           .word x010,x011,x018,x019
           .word x100,x101,x108,x109
           .word x110,x111,x118,x119
           .word x800,x801,x808,x809
           .word x810,x811,x818,x819
           .word x900,x901,x908,x909
           .word x910,x911,x918,x919


x000
x008
x800
x808
           .block
           lda i4
           sta r4
           lda #$00
           sta rd
           lda be
           sta re
           jmp compare
           .bend


x001
x801
           .block
           lda i4
           sec
           sbc #2
           ldx i4
           cpx #3
           bcs noload
           lda b4
           cpx #0
           bne nodec
           cmp #2
           bcc nodec
           sec
           sbc #1
nodec
noload
           sta r4
           lda #$00
           ldx i4
           cpx #7
           bcs nobit0
           ora #$01
nobit0
           cpx #6
           bcs nobit7
           ora #$80
nobit7
           sta rd
           lda #$01
           sta re
           jmp compare
           .bend


x009
x809
           .block
           lda i4
           sec
           sbc #2
           ldx i4
           cpx #3
           bcs noload
           lda b4
noload
           sta r4
           lda #$00
           ldx i4
           cpx #7
           bcs nobit0
           ora #$01
nobit0
           cpx #6
           bcs nobit7
           ora #$80
nobit7
           sta rd
           lda #$09
           ldx i4
           cpx #$0b
           bcs nostop
           and #$08
nostop
           sta re
           jmp compare
           .bend


x010
x018
x810
x818
           .block
           lda b4
           sta r4
           lda #$00
           sta rd
           lda be
           and #$09
           sta re
           jmp compare
           .bend


x011
x811
           .block
           ldx b4
           cpx #2
           bcc nodec
           dex
nodec
           stx r4
           lda #$00
           ldx b4
           cpx #6
           bcs nobit0
           ora #$01
nobit0
           cpx #5
           bcs nobit7
           ora #$80
nobit7
           ldx i4
           bne nobit07
           ora #$81
nobit07
           sta rd
           sta rd
           lda be
           and #$09
           sta re
           jmp compare
           .bend


x019
x819
           .block
           ldx b4
           cpx #2
           bcc nodec
           lda i4
           beq nodec
           dex
nodec
           stx r4
           lda #$00
           ldx b4
           cpx #6
           bcs nobit0
           ora #$01
nobit0
           cpx #5
           bcs nobit7
           ora #$80
nobit7
           ldx i4
           bne nobit07
           ora #$81
nobit07
           sta rd
           sta rd
           lda #$09
           ldx i4
           beq stop
           ldx b4
           cpx #$0a
           bcs nostop
stop
           and #$08
nostop
           sta re
           jmp compare
           .bend


x100
x108
           .block
           lda i4
           ldx #$00
           sec
           sbc #$0b
           bcs noload
           lda b4
           ldx i4
           cpx #$0a
           bcs nosub
           sec
           sbc subtab,x
           bcs nosub
           lda i4
           asl a
           asl a
           asl a
           asl a
           ora b4
           ldx #corr-special-1
search
           cmp special,x
           beq found
           dex
           bpl search
           lda b4
           jmp nosub
found
           lda corr,x
nosub
           ldx #$81
noload
           sta r4
           stx rd
           lda be
           sta re
           jmp compare
subtab     .byte 5,5,5,3,1,5,4,3,2,1
special    .byte $71,$62,$53,$52,$51
           .byte $31,$23,$22,$21,$13
           .byte $12,$11,$03,$02,$01
           .byte $00
corr       .byte $00,$01,$02,$00,$00
           .byte $00,$02,$00,$00,$02
           .byte $00,$00,$02,$00,$00
           .byte $00
           .bend


x101
           .block
           lda i4
           sec
           sbc #$0d
           beq load81
           bcc load81
           cmp #$04
           bcc set81
           beq set01
           ldx #$00
           jmp set
set01
           ldx #$01
           jmp set
load81
           lda b4
           ldx i4
           cpx #$0c
           bcs set81
           sec
           sbc subtab,x
           beq test
           bcs set81
test
           lda i4
           asl a
           asl a
           asl a
           asl a
           ora b4
           ldx #corr-special-1
search
           cmp special,x
           beq found
           dex
           bpl search
           lda b4
           jmp set81
found
           lda corr,x
set81
           ldx #$81
set
           sta r4
           stx rd
           lda #$01
           sta re
           jmp compare
subtab     .byte 7,7,7,5,3,7,6,5,4,3,2
           .byte 1
special    .byte $82,$73,$64,$63,$55
           .byte $54,$52,$33,$25,$24
           .byte $22,$15,$14,$12,$05
           .byte $04,$02
corr       .byte $01,$02,$03,$01,$04
           .byte $02,$01,$02,$04,$02
           .byte $01,$04,$02,$01,$04
           .byte $02,$01
           .bend


x109
           .block
           lda i4
           sec
           sbc #$0d
           beq load81
           bcc load81
           cmp #$04
           bcc set81
           beq set01
           ldx #$00
           jmp set
set01
           ldx #$01
           jmp set
load81
           lda b4
           ldx i4
           cpx #$0c
           bcs set81
           sec
           sbc subtab,x
           beq reload
           bcs set81
reload
           lda b4
set81
           ldx #$81
set
           sta r4
           stx rd
           ldy #$08
           ldx i4
           cpx #$16
           bcs start
           cpx #$0a
           bcs sete
           lda b4
           cmp b4comp,x
           bcc sete
start
           ldy #$09
sete
           sty re
           jmp compare
subtab     .byte 7,7,7,5,3,7,6,5,4,3,0
           .byte 0
b4comp     .byte $10,$10,$10,$0e,$0c
           .byte $10,$0f,$0e,$0d,$0c
           .bend


x110
x118
           .block
           lda b4
           sta r4
           lda #$00
           ldx i4
           cpx #$0b
           bcs nofire
           lda #$81
nofire
           sta rd
           lda be
           and #$09
           sta re
           jmp compare
           .bend


x111
           .block
           ldx b4
           cpx #2
           bcc nodec
           dex
nodec
           stx r4
           lda #$00
           ldx b4
           cpx #6
           bcs nobit0
           ora #$01
nobit0
           cpx #5
           bcs nobit7
           ora #$80
nobit7
           ldx i4
           bne nobit07
           ora #$81
nobit07
           ldx i4
           cpx #$0c
           bcs nofire
           lda #$81
nofire
           sta rd
           lda be
           and #$09
           sta re
           jmp compare
           .bend


x119
           .block
           ldx b4
           cpx #2
           bcc nodec
           lda i4
           cmp #$0c
           bcs dodec
           cmp #$0a
           bcs nodec
           cpx #$0f
           bcs dodec
           asl a
           asl a
           asl a
           asl a
           ora b4
           ldy #$12
search
           cmp nodectab,y
           beq nodec
           dey
           bpl search
dodec
           dex
nodec
           stx r4
           lda #$00
           ldx b4
           cpx #6
           bcs nobit0
           ora #$01
nobit0
           cpx #5
           bcs nobit7
           ora #$80
nobit7
           ldx i4
           cpx #$0c
           bcs nobit07
           ora #$81
nobit07
           sta rd
           lda #$09
           ldx i4
           cpx #$0a
           bcc teststop
           cpx #$0c
           bcc stop
teststop
           ldy b4
           cpy #$0a
           bcs nostop
stop
           lda #$08
nostop
           sta re
           jmp compare
nodectab   .byte $82,$73,$72,$64,$63
           .byte $55,$54,$52,$33,$32
           .byte $25,$24,$22,$15,$14
           .byte $12,$05,$04,$02
           .bend


x900
x908
           .block
           lda i4
           cmp #$05
           bcc set81
           ldx #$00
           sec
           sbc #$0b
           bcs noload
           lda b4
set81
           ldx #$81
noload
           sta r4
           stx rd
           lda be
           sta re
           jmp compare
           .bend


x901
           .block
           lda i4
           cmp #$0e
           bcs subd
           cmp #$04
           beq sub2
           cmp #$03
           beq sub2
           tax
           lda b4
           sec
           sbc subtab,x
           beq load
           bcs set4
load
           lda b4
           jmp set4
subd
           sec
           sbc #$0d
           jmp set4
sub2
           sec
           sbc #$02
set4
           sta r4

           ldx #$00
           lda i4
           cmp #$11
           bne nobit0
           ldx #$01
nobit0
           bcs nobit7
           ldx #$81
nobit7
           stx rd
           lda #$01
           ldx i4
           cpx #$0a
           bne nostop
           lda #$00
nostop
           sta re
           jmp compare
subtab     .byte 1,0,0,0,0,2,2,2,2,2
           .byte 0,1,0,0
           .bend


x909
           .block
           lda i4
           cmp #4
           beq sub2
           cmp #3
           beq sub2
           sec
           sbc #$0d
           beq load
           bcs noload
load
           ldx i4
           lda b4
           sec
           sbc subtab,x
           beq reload
           bcs noload
reload
           lda b4
           jmp noload
sub2
           sec
           sbc #2
noload
           sta r4
           lda #$00
           ldx i4
           cpx #$11
           bne nobit0
           lda #$01
nobit0
           bcs nobit7
           lda #$81
nobit7
           sta rd
           lda #$08
           ldx i4
           cpx #$16
           bcs start
           cpx #$0a
           bcs sete
           cpx #$05
           bcc sete
           ldx b4
           cpx #$0b
           bcc sete
start
           lda #$09
sete
           sta re
           jmp compare
subtab     .byte 0,0,0,0,0,2,2,2,2,2
           .byte 0,0,0,0
           .bend


x910
x918
           .block
           lda b4
           sta r4
           lda #$00
           ldx i4
           cpx #$0b
           bcs setd
           lda #$81
setd
           sta rd
           lda be
           and #$09
           sta re
           jmp compare
           .bend


x911
           .block
           lda b4
           ldx i4
           cpx #$0a
           beq noload
           sec
           sbc #$01
           beq load
           bcs noload
load
           lda b4
noload
           sta r4
           lda i4
           cmp #$0c
           bcc set81
           lda #$00
           ldx b4
           cpx #$05
           bne nobit0
           lda #$01
nobit0
           bcs nobit7
set81
           lda #$81
nobit7
           sta rd
           lda be
           and #$09
           ldx i4
           cpx #$0a
           bne nostop
           lda #$00
nostop
           sta re
           jmp compare
           .bend


x919
           .block
           ldx i4
           beq load
           cpx #$0b
           beq load
           cpx #$0a
           beq load
           lda b4
           sec
           sbc #$01
           beq load
           bcs noload
load
           lda b4
noload
           sta r4
           ldx i4
           cpx #$0c
           bcc set81
           lda #$00
           ldx b4
           cpx #$05
           bne nobit0
           lda #$01
nobit0
           bcs nobit7
set81
           lda #$81
nobit7
           sta rd
           lda #$09
           ldx i4
           beq stop
           cpx #$0a
           bcc testb
           cpx #$0c
           bcc stop
testb
           ldx b4
           cpx #$0a
           bcs nostop
stop
           lda #$08
nostop
           sta re
           jmp compare
           .bend


compare
           lda a4
           cmp r4
           bne error
           lda ad
           cmp rd
           bne error
           lda ae
           cmp re
           bne error
noerror
           inc beindex
           lda beindex
           cmp #8
           bcc jmploop
           lda #0
           sta beindex
           inc ieindex
           lda ieindex
           cmp #4
           bcc jmploop
           lda #0
           sta ieindex
           dec b4
           bpl jmploop
           lda #20
           sta b4
           dec i4
           bpl jmploop
           jmp finish
jmploop
           jmp loop
error
           jsr print
           .byte 13,13
           .text "init  "
           .byte 0
           lda i4
           jsr printhb
           lda #32
           jsr $ffd2
           lda b4
           jsr printhb
           lda #32
           jsr $ffd2
           lda ie
           jsr printhb
           lda #32
           jsr $ffd2
           lda be
           jsr printhb
           jsr print
           .byte 13
           .text "after "
           .byte 0
           lda a4
           jsr printhb
           lda #32
           jsr $ffd2
           lda ad
           jsr printhb
           lda #32
           jsr $ffd2
           lda ae
           jsr printhb
           jsr print
           .byte 13
           .text "right "
           .byte 0
           lda r4
           jsr printhb
           lda #32
           jsr $ffd2
           lda rd
           jsr printhb
           lda #32
           jsr $ffd2
           lda re
           jsr printhb
           jsr waitkey
           jmp noerror
finish
           .bend


           jsr print
           .text " - ok"
           .byte 13,0
           lda turboass
           beq load
           jsr waitkey
           jmp $8000
load
           jsr print
name       .text "cia1tb"
namelen    = *-name
           .byte 0
           lda #0
           sta $0a
           sta $b9
           lda #namelen
           sta $b7
           lda #<name
           sta $bb
           lda #>name
           sta $bc
           pla
           pla
           jmp $e16f


