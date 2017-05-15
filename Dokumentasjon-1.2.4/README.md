# Oppgave 1.2.4

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

Beskrivelse av koden log.go: Dette er en pakke som har blitt kalt «function». Koden inneholder import av fmt, os, math og strconv.
Pakken fmt gir oss flere muligheter til å formatere ulike verdier, dette gjøres ved hjelp av utskrift «verb».
OS pakken gir en plattform som er uavhengig av grensesnitt til operativ-funksjonalitet.
Pakken «math» gir grunnleggende konstanter og matematiske funksjoner, «strconv» gir muligheten til å konvertere til og fra «String» av grunnleggende data typer.

På linje 10 linje har vi en funksjon med navet «LogariteFunc». "Base" står for argument 1 og "tall" står for argument 2. Tall float på linje 13 konverterer argument 2 til en float 64 verdi.
På linje 14 konverters argument 1 fra string til int ved hjelp av «.Atoi».
Beskrivelse av koden main.go: I denne filen har vi main pakken, som importere pakken function. Og kjører funksjonen som ligger i den.
Mot slutten ser vi en if-setning som sier: hvis basefloat er det samme som 2, så printer den ut svare i form av math.log2, dette betyr at den blir beregnet med base av 2.
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
