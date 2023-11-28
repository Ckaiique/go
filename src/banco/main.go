package main

import ("fmt")


type ContaCorrente struct {
	titular string
	numeroAgecnia int
	numeroConta int
	saldo float64
}

func mains() {
	// titular := "Kaique"
	// numeroAgecnia := 589
	// numeroConta := 123456
	// saldo := 125.5

	//fmt.Println(titular,numeroAgecnia,numeroConta,saldo)
	contaDoKaique  := ContaCorrente{
		titular: "kaique",
		numeroAgecnia: 589,
		numeroConta: 123456,
		saldo: 125.5,
	}
	fmt.Println(contaDoKaique)

	contaDaLeticia := ContaCorrente{"Letica",222,111222,200}
	fmt.Println(contaDaLeticia)
	
}
