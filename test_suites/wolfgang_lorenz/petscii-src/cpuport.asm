           *= $0801           .BYTE $4C,$14,$08,$00,$97TURBOASS   = 780           .TEXT "780"           .BYTE $2C,$30,$3A,$9E,$32,$30           .BYTE $37,$33,$00,$00,$00           LDA #1           STA TURBOASS           JMP MAINCONFIG     .BYTE 0ABACKUP    .BYTE 0,0LASTSTATE  .BYTE 0RIGHT      .BYTE 0ROM           LDA #$2F           STA 0           LDA #$37           STA 1           CLI           RTSMAIN           JSR PRINT           .BYTE 13           .TEXT "�CPUPORT"           .BYTE 0           LDA #0           STA CONFIGNEXTCONFIG           SEI           LDA #$FF           STA 0           STA 1           STA ABACKUP+0           STA ABACKUP+1           STA LASTSTATE           LDX #8           LDA CONFIGPUSH           ASL A           PHP           DEX           BNE PUSH           LDX #4PULL           PLA           AND #1           TAY           LDA #0           PLP           SBC #0           STA 0,Y           STA ABACKUP,Y;INPUTS: KEEP LAST STATE           LDA ABACKUP+0           EOR #$FF           AND LASTSTATE           STA OR1+1;OUTPUTS: SET NEW STATE           LDA ABACKUP+0           AND ABACKUP+1OR1        ORA #$11           STA LASTSTATE;DELAY FOR LARGER CAPACITIVES           LDY #0DELAY           INY           BNE DELAY           DEX           BNE PULL           LDA ABACKUP+0           CMP 0           BNE ERROR           LDA ABACKUP+0           EOR #$FF           ORA ABACKUP+1           AND #$37           STA OR2+1           LDA LASTSTATE           AND #$C8OR2        ORA #$11;BIT 5 IS DRAWN LOW IF INPUT           TAX           LDA #$20           BIT ABACKUP+0           BNE NO5LOW           TXA           AND #$DF           TAXNO5LOW           STX RIGHT           CPX 1           BNE ERRORNOERROR           INC CONFIG           BNE NEXTCONFIG           JSR ROM           JMP OKERROR           LDA 1           PHA           LDA 0           PHA           JSR ROM           JSR PRINT           .BYTE 13           .TEXT "0=FF 1=FF"           .BYTE 0           LDX #8           LDA CONFIGPUSH1           ASL A           PHP           DEX           BNE PUSH1           LDX #4PULL1           LDA #32           JSR $FFD2           PLA           AND #1           ORA #"0"           JSR $FFD2           LDA #"="           JSR $FFD2           LDA #0           PLP           SBC #0           STX OLDX+1           JSR PRINTHBOLDX           LDX #$11           DEX           BNE PULL1           JSR PRINT           .BYTE 13           .TEXT "AFTER  "           .BYTE 0           PLA           JSR PRINTHB           LDA #32           JSR $FFD2           PLA           JSR PRINTHB           JSR PRINT           .BYTE 13           .TEXT "RIGHT  "           .BYTE 0           LDA ABACKUP+0           JSR PRINTHB           LDA #32           JSR $FFD2           LDA RIGHT           JSR PRINTHB           LDA #13           JSR $FFD2WAITK           JSR $FFE4           BEQ WAITK           CMP #3           BEQ STOP           JMP NOERRORSTOP           LDA TURBOASS           BEQ BASIC           JMP $8000BASIC           JMP $A474OK           JSR PRINT           .TEXT " - OK"           .BYTE 13,0           LDA TURBOASS           BEQ LOADWAIT       JSR $FFE4           BEQ WAIT           JMP $8000LOAD           LDA #47           STA 0           JSR PRINTNAME       .TEXT "CPUTIMING"NAMELEN    = *-NAME           .BYTE 0           LDA #0           STA $0A           STA $B9           LDA #NAMELEN           STA $B7           LDA #<NAME           STA $BB           LDA #>NAME           STA $BC           PLA           PLA           JMP $E16FPRINT      PLA           .BLOCK           STA PRINT0+1           PLA           STA PRINT0+2           LDX #1PRINT0     LDA !*,X           BEQ PRINT1           JSR $FFD2           INX           BNE PRINT0PRINT1     SEC           TXA           ADC PRINT0+1           STA PRINT2+1           LDA #0           ADC PRINT0+2           STA PRINT2+2PRINT2     JMP !*           .BENDPRINTHB           .BLOCK           PHA           LSR A           LSR A           LSR A           LSR A           JSR PRINTHN           PLA           AND #$0FPRINTHN           ORA #$30           CMP #$3A           BCC PRINTHN0           ADC #6PRINTHN0           JSR $FFD2           RTS           .BEND