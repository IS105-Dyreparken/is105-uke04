# Oppgave 1.2.2

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
