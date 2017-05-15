# Oppgave 1.2.1

De fleste kjenner til 10-tallsystemet(desimaltallsystemet), dette systemet er basert på det.
Det finnes flere tallsystemer. Det binære(2-tallsystemet) er et av disse, som er da basert på 2 tall(0 og 1).
Det er dette tallsystemet som maskinen forstår, og det er derfor nødvendig for oss å kjenne til bruken av det.
For å telle binære tall leses disse fra høyre til venstre, i motsetning til 10-tallsystemet blir binære tall-verdier 2 ganger så høye for hver sifferplassering til venstre.
Et eksempel på hvordan konvertering av 2 tallsystemet funker

|   |   |   |   |   |
|---|---|---|---|---|
|2^4|2^3|2^2|2^1|2^1|
|16 |8  |4  |2  |1  |
|0  |1  |0  |1  |1  |


For å regne ut tall der du ikke finner verdien direkte i tallsystemet, slik som tallet 5, tar du det binære tallet som er nærmest; f.eks med tallet 22:22 = 16 + 6. Problemet som oppstår her er at tallet 6 finnes ikke i 2-tallsystemet, for å løse dette problemet kan vi gjøre følgende.

22 = 16 + 4 + 2 (Da har vi funnet tallene i totallsystemet, og setter de sammen til binær kode)
22 = 2^4 + 2^2 + 2^1
22 = 10110
Når vi jobber med både 2-tallsystemet og 10-tallsystemet er det vanlig å skrive bak tallene hvilket system de tilhøre eks. 10110(to) = 22(ti).

(referanse fra: https://github.com/Zwirc/IS-105/tree/master/ICA01)

Konvertering av desimaltall til 2-tallsytemet.
Oppgaven går ut på å regne fra desimal til binær, da må vi regne fra høyeste tall til laveste.
(1) 1		1(ti) = 2^0 dette blir da 1bit som tilsvare 1 binært
(2) 2		2(ti) = 2^1 dette blir da 2bit som tilsvare 10 binært
(3) 5		5(ti) = 2^2 + 2^0 dette blir da 3bit som tilsvare 101 binært
(4) 8		8(ti) = 2^3 dette blir da 4bit som tilsvarer 1000 binært
(5) 16		16(ti) = 2^4 dette blir da 5bit som tilsvarer 10000 binært
(6) 256		256(ti) = 2^8 dette blir da 9 bit som tilsvarer 100000000 binært

Konverter binært til desimal
Når vi regner fra binært til desimal, så legge vi sammen alle verdier som er 1
(1) 100(to)		2^2 = 4(ti)

|   |   |   |
|---|---|---|
|2^2|2^1|2^0|
|1  |0  |0  |
|0  |1  |1  |

4+0+0 = 4(ti)


(2) 1001(to)	2^3 + 2^0 = 9(ti)

|   |   |   |   |
|---|---|---|---|
|2^3|2^2|2^1|2^0|
|1  |0  |0  |1  |

8+0+0+1=9(ti)


(3) 1100110011(to)  2^9 + 2^8 + 2^5 +2^4 +2^1 + 2^0 = 819(ti)

|   |   |   |   |   |   |   |   |   |   |
|---|---|---|---|---|---|---|---|---|---|
|2^9|2^8|2^7|2^6|2^5|2^4|2^3|2^2|2^1|2^1|
|1  |1  |0  |0  |1  |1  |0  |0  |1  |1  |

512+ 256 +0+0+32+16+0+0+2+1 = 819(ti)


kilde: https://www.diskusjon.no/index.php?showtopic=765979
