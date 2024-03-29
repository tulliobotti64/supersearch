package main

import (
	"fmt"
	"strings"
)

func main() {
	textHtml := `<p>I font e i typeface sono termini spesso usati in modo intercambiabile, ma in realtà ci sono delle differenze importanti tra i due. In questo articolo esploreremo cosa sono i font e i typeface, le principali differenze tra di loro e come riconoscere i vari tipi di typeface. Inoltre, vedremo perché è importante scegliere il giusto typeface per un progetto.</p>
<h3>Cosa sono i font e i typeface?</h3>
<p>I font sono un insieme di caratteri tipografici progettati per essere utilizzati in un determinato formato, come il carattere Times New Roman per il formato di testo di un documento word o Arial per il formato web. I font contengono le <strong>informazioni su come il carattere deve essere visualizzato</strong>, come la dimensione, lo spessore del tratto, l'inclinazione e così via.</p>
<p>I typeface, invece, sono un insieme di caratteri tipografici progettati per essere utilizzati insieme per creare una <strong>famiglia di font</strong>. Ad esempio, Times New Roman è un typeface che include diverse varianti come Times New Roman Regular, Times New Roman Bold e Times New Roman Italic.</p>
<blockquote>
<p>“La principale differenza tra font e typeface è che i font sono un singolo set di caratteri tipografici, mentre i typeface sono un insieme di font che condividono caratteristiche stilistiche simili.”</p>
</blockquote>
<h3>Come riconoscere i vari tipi di font:</h3>
<p>I font possono essere suddivisi in diverse categorie in base alle loro caratteristiche:</p>
<ul>
<li>I font slab serif sono caratterizzati da &quot;piedini&quot; più spessi rispetto ai font serif tradizionali. Sono spesso utilizzati per i progetti di design, come i loghi e i titoli.
<img src="/storage/media/153fbcaf817d3b276280ea20fb5fa16b/slab serif sheriff.jpg" alt="slab-serif.jpg" /></li>
<li>I font variable sono una nuova tipologia di font che permette di modificare la dimensione, lo spessore e l'inclinazione dei caratteri in modo dinamico. Questo rende possibile creare effetti visivi unici e adattare il font alle esigenze del progetto.
<img src="/storage/media/153fbcaf817d3b276280ea20fb5fa16b/font variable.jpeg" alt="font-variable.jpeg" /></li>
<li>I font serif sono caratterizzati da piccole linee o &quot;piedini&quot; alla fine dei tratti dei caratteri. Sono spesso utilizzati per i testi scritti a mano o stampati, come i libri e i giornali.
<img src="/storage/media/153fbcaf817d3b276280ea20fb5fa16b/serif 2.webp" alt="serif.webp" /></li>
<li>I font sans-serif non hanno &quot;piedini&quot; alla fine dei tratti dei caratteri. Sono spesso utilizzati per i testi digitali, come i siti web e le interfacce utente.
<img src="/storage/media/153fbcaf817d3b276280ea20fb5fa16b/sans serif.webp" alt="sans-serif.webp" /></li>
<li>I font script sono progettati per imitare la scrittura a mano e sono spesso utilizzati per i progetti artistici e creativi.
153fbcaf817d3b276280ea20fb5fa16b/script.jpeg" alt="script.jpeg" /></li>`
	//result : ""
	var arrText []string
	var searchIndex int = 0
	searchArr := `<img src="/storage/media/` //24
	searchArrEnd := `"`
	var achou bool
	var oneChar string
	for x := 0; x < len(textHtml); x++ {
		oneChar = textHtml[x : x+1]
		if achou {
			if oneChar == searchArrEnd {
				achou = false
				searchIndex = 0
			} else {
				if oneChar == " " {
					oneChar = "_"
				}
			}
		}
		arrText = append(arrText, oneChar)

		if !achou {
			if oneChar == searchArr[searchIndex:searchIndex+1] {
				if len(searchArr) == searchIndex+1 {
					achou = true
				}
				searchIndex++
			} else {
				searchIndex = 0
			}
		}
	}
	fmt.Println(strings.Join(arrText, ""))
}
