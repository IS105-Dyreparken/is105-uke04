# is105-uke04 Dokumentasjon/besvarelse av oppgavene
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


#Oppgave 1.2.2

Informasjonsmengde Flere personer prøver å gjette et tresifret (3-bit) binært tall.

(1) Lise har fått vite / lærer at tallet er et oddetall.
(2) Per har fått vite at tallet er IKKE et multiplum av 3 (dvs. ikke 0, 3, 6).
(3) Oskar har fått vite at tallet inneholder nøyaktig 2 enere.
(4) Louise har fått vite alt det Lise, Per og Oskar vet.

|   |   |   |   |   |   |   |   |
|---|---|---|---|---|---|---|---|
|7  |6  |5  |4  |3  |2  |1  |0  |
|111|110|101|100|011|010|001|000|

Dette er muligheten vi har når det er et binærtall på 3 bit.

Lise har fått vite at det er oddetall ut i fra tabellen har Lise disse tallene:

|   |   |   |   |
|---|---|---|---|
|7  |5  |3  |1  |
|111|101|011|001|


Per har fått vite at tallene ikke er et multiplum av 3, det vil si at tallet ikke kan være 3 eller gange 3.
Da gjenstår disse tallene fra tabellen:

|   |   |   |   |   |
|---|---|---|---|---|
|7  |5  |4  |2  |1  |
|111|101|100|010|001|



Oskar får vite at tallet inneholder nøyaktig 2 enere, ut ifra dette vet vi at det er snakk om de binære tallene. Dette er da mulighetene han har:

|   |   |   |
|---|---|---|
|6  |5  |3  |
|110|101|011|



Louise har fått vite alt det de andre vet. Ut i fra tabellene som er opprettet ut ifra deres tilgjengelige informasjon leter vi etter det tallet som gjentar seg i alle 3 tabellene. Ut i fra tabellen kan Louise se at det eneste tallet som er felles for alle tabellene er tallet 5(ti) 101(to).

Hvor mye informasjon i bits har hver person fått vite?  
Vi har totalt 8 muligheter som er 3 sifferet binære tall. Vi kaller variabelen for muligheten for X, og vi får dermed X = 8. Muligheten i bits blir da Log2(X/muligheter) = bits. Mengde informasjon som denne oppgaven har i bits blir da Log2(8) = 3 bits
1.	Lise har fått vite at tallet er et oddetall, dette gir henne 4 muligheter og tilsvare en informasjonsmengde på: Log2(8/4) = 1bit

2.	Per har fått vite at tallet er ikke et multiplum av 3 han har da 8-3 = 5 muligheter, dette tilsvare en informasjonsmengde på: Log2(8/5) = 0 bit

3.	Oskar har fått vite at tallet inneholder nøyaktig 2 enere han har da 8-5 = 3muligheter, dette tilsvare en informasjonsmengde på: Log2(8/3) = 1bit

4.	Louise har fått vite alt det de andre har fått vite hun har da 8-7 = 1mulighet som passer beskrivelsene, dette tilsvare en informasjonsmengde på: Log2(8/1) = 3bits


#Oppgave. 1.2.3

Opprettelse av en ny git repository med navnet “is105-uke04” på Github med README, gitignore, og lisens fil er gjort på følgende plass: https://github.com/IS105-Dyreparken/is105-uke04
Fremgangsmetoden vi brukte for å klone repository på våre systemer:
1.	Åpne en kommandolinje på systemet vårt(vi som har Windows operativsystem brukte «windows power shell»)

2.	Naviger deg til den mappen dere ønsker å lagre repository i. Navigasjonen ble gjort ved hjelp av funksjonen cd «mappenavn», og dannelse av mappen ble gjort med funksjonen mkdir.
Eksempel: mkdir ICA01(lager mappen med navnet ICA01).	cd ICA01(navigere deg fra gjeldene plassering til mappen ICA01).

3.	Klone repository ved å kopiere adressen til git repository og deretter skrive følgende i shell:
git clone https://github.com/IS105-Dyreparken/is105-uke04.

4.	Opprettelse av en egen branch i repository med vårt eget navn. Dette gjøres ved å skrive i terminal: git checkout -b «Navn»

5.	Opprett en ny go fil i mappen via kommandolinje. Når du har lagret filen kan du gå inn og endre den ved hjelp av en editor som f.eks vim. Skriver vi «git status» i terminal kan vi se at filen er rød og det betyr at det er gjort en endring.

6.	 For å klargjøre filen for opplastning skriver vi «git add .» som gjør at du legger til alle filer.
 git commit -m «»kommandoen «»gjør at du kan kommentere på endringen.

7.	Git push origin «branch navn» for laste opp endringer til branch.

#Oppgave 1.2.4

1. Fordeler og ulemper med git-flow-modell i en hovedrepository
Fordeler: Lett tilgjengelig for dem som jobber med koden, og koden blir lett å dele ved å sette tilgangen «public». Mange kan jobbe med forskjellige deler av koden, i hver sin kopi av hovedrepository. Dette fører til at mer kan bli gjort samtidig og alle endringer kan sammenlignes før det evt. lastes opp til hovedrepository.
Ulemper: Siden alle har tilgang til koden og alle kan gjøre endringer samtidig, kan dette føre til at filer blir slettet. Endringer blir gjort i «master-branchen» i forhold til «test-branchen». Siden det er så mange som kan jobbe på en repository samtidig kan det fort bli kaotisk, ved at ting blir opprettet uten dokumentasjon eller at dokumentasjon mangler generelt.

2.	Objektfiler for de mest brukte plattformer (Unix/Linux, MS Windows, Mac OS X) Hvorfor, etter deres mening, har disse plattformene så forskjellige objektfil-formater
Windows - EXE
Mac OS X - Mach-O
Linux – ELF
Disse plattformene har forskjellige objekt filer siden de er strukturert på forskjellige måter.
OS X, Windows og Unix/Linux har forskjellige filstrukturer, og kjører på forskjellige kernels.
En fil designet for å tildele verdier til registre i Windows ville derfor fungert dårlig på en Unix plattform.

3.	Forskjeller mellom java og golang
Det er forskjeller fra java, blant annet ved import statementet, hvor man kan importere slik som java, eller flere pakker samtidig med
Import (
Pakke1
Pakke2)
Man trenger også generelt mindre kode i Golang, da den automatisk skjønner mye
som må manuelt defineres i java (som for eksempel break i en switch, dette tar golang automatisk).
Ellers så slipper man semikolon på enden av hver setning
I if statements og lignende slipper man også å bruke () for å holde statementet

Kode for oppgave 4 og 5 er hentet fra: https://github.com/Gruppe12IS105/ICA01.
4.	Hvilke viktige poeng illustrerer denne øvelsen når det gjelder bruk av et programmeringsmiljø på en plattform?
For å kjøre main.go filen må man skrive: «go run main.go x y», der x og y er tall. X må være enten tallet 2 eller 10, Y kan være hvilket som helst tall.
Filen log.go innholder koden for deloppgave 6 også det er grunnen for at man må definere både X og Y- verdi.

Beskrivelse av koden log.go: Dette er en pakke som har blitt kalt «function». Koden inneholder import av fmt, os, math og strconv. Pakken fmt gir oss flere muligheter til å formatere ulike verdier, dette gjøres ved hjelp av utskrift «verb». OS pakken gir en plattform som er uavhengig av grensesnitt til operativ-funksjonalitet. Pakken «math» gir grunnleggende konstanter og matematiske funksjoner, «strconv» gir muligheten til å konvertere til og fra «String» av grunnleggende data typer.
På linje 10 linje har vi en funksjon med navet «LogariteFunc». Base står for argument 1 og tall står for argument 2. Tall float på linje 13 konverterer argument 2 til en float 64 verdi. På linje 14 konverters argument 1 fra string til int ved hjelp av «.Atoi».
Beskrivelse av koden main.go: I denne filen har vi main pakken, som importere pakken function. Og kjører funksjonen som ligger i den.
Til slutt har vi en if setning som sier: hvis basefloat er det samme som 2, så printer den ut svare i form av math.log2, dette betyr at den blir beregnet med base av 2.
Hvis basefloat er 10 blir det samme gjort, men svaret blir beregnet i form av math.log10 som er beregning gjort i base 10.
Til slutt vil vi få en veiledende respons om kriteriene for at kjøring av filen ikke er oppfylt.  
Alle programmeringsmiljøer er forskjellige, og krever derfor litt forskjellig tankegang.
I Go sitt tilfelle er riktig mappestruktur en sentral del for at alt skal fungere slik man vil.
Hvis man setter GOPATH feil, eller ikke legger prosjektene i GOPATH\src, så vil ikke programmene kjøre.

5.	Er det hensiktsmessig å legge inn denne filen i git repository? Begrunn svaret!

Beskrivelse av filen logcli: Denne filen er lik som log.go, utenom at den er en main pakke og har kun et argument. Denne filen printer også ut kun logaritmen av et tall i base 2, hvordan dette skjer spesifikt er det samme som i forklaring i oppgaven ovenfor.

Det kan være hensiktsmessig å legge denne i et git repository,
man har da backup av filen, og i dette tilfellet hvor den er en del av en innlevering
så gir det mening å legge den sammen med de andre delene av innleveringen.


6.	Hva skiller pakke log, fra andre pakker i go?
log pakken som vi har implementert skiller seg ut ved at den er en liten pakke, for eksempel i motsetning til fmt.
Log har også kun en funksjon, som betyr at den kun gjør en liten ting, mens fmt har mange bruksområder
eks Println.
