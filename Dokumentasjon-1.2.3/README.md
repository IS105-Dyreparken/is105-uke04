# Oppgave. 1.2.3

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
