package main

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
)

func main() {
	var buf bytes.Buffer
	source := `
		# Quae Procrusten putatur terraeque aures mirum inhospita
		
		## Cremat naufraga clarique socios abscessisse obtulimus pertimui
		
		Lorem markdownum votis, ut arma tenero nondum obscura Nise, undaeque ad reicere
		cornua Memnonis, ego, in? E tu serasque *stabat*! Tamen
		[pia](http://memorant.com/deposcunt) mater.
		
		> Vulnere adeste, lapis coniunxque amnis viridique solumque ducum sed; ita
		> repellit; loca bellum. Propiore adspicit equo locis, et herbida Aeneas,
		> miracula *et dicta* pendentia primis in adfecit petite. Milon pro surgit, atro
		> quae. Licet tuum pedes solibus.
		
		**Ponat ore** tuaeque velox dum fatis, quod dum vulnere fertur noctisque.
		**Similisque unda ponto** usurus [radiare Amazone
		linguae](http://remos-haud.net/), cum nec multa vix, et conscia dentes
		adspiceres lymphae. Sine nec saevit opprobria, tunc qui infra dissipat tecum,
		flammas, somnum repetita **et pudetque aperti**. Mediam in *o* sua, O radice
		crurorem deducat longo demersit quoque tenet, in **non**.
		
		## Amnes in manu
		
		In Erycis famulae puer onus apicemque portante fluit! Ipsi et Mavortis
		aquilonibus dies pressa Liber fusile.
		
		1. Placidique perdant perque
		2. Vindexque flosque
		3. Studiis iamque iamque lacrimis agris Thisbe
		
		Dextra Sicanias ad, vult **fraterno**, subiectas ad nunc plumbo Acmon et ad
		partem furenti forsitan [preces](http://dat-tela.net/) ultor. Et inflataque
		pietasque melior; serta subiectum *relevat semper supernum* Acesten adeunt,
		sceptrum de **tacet**. Orbem et nocent quoque vittisconcutiens aestuat
		descenderat neque parte summo, silent lacrimoso.
		
		## Id diu vel iram poteris coxerat nec
		
		Mediisque clamoribus, amara nec et ista per utrumque fata. Illa nemus: crine
		manu soror delphines aliis sacrata oculosque inposito sub inter et **procul
		haliaeetus**. Repagula fama pomis aures, dura alios trahentem Martem dimovit,
		de? Aut quoque aquis vesci et stipulis quacumque [permulcet
		cornu](http://crathis.io/me-vultus.html), sic nec per exiguus.
		
		> Murmura lacteus possem. Praedator ut umero, et fontem et mutati denique sibi
		> nisi: in adurat! Laticem quantam quod [quam](http://www.non.net/) recolligat
		> suae: *tuorum ecce*. Videt erit; aequos animae victrixque, age exstat
		> auxiliaribus adparuit?
		
		Obsidis gemino Thybris e ieiunia voce totiens Lucifer inveniunt fuit. Scelerata
		scitusque auras molimine, funera et origine caede. [Capillo facit
		plura](http://partes.io/frustra) fervens penatigero umero subdit ducit amori
		maxima relicto orbem Olympi inter veniens Circaea, hospes tibi letifer! Pressit
		lacrimas alias gratantur famulosque eras te deus cognoscere nisi cernimus,
		timeas tria trabibus, hederarum.
	`
	if err := goldmark.Convert([]byte(source), &buf); err != nil {
		panic(err)
	}
	result := buf.String()
	fmt.Println(result)
}
