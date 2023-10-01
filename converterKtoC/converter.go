package main
import "fmt"

func main() {
  const ebulicaoAguaK int16 = 373 //Graus Celcius
  var Celcius int16 = 0
  var Kelvin int16 = Celcius + 273

  var ebulicaoAguaC int16 = ebulicaoAguaK - Kelvin

  fmt.Printf("In Kelvin, the boiling point of Water(H2O) is 373K\n"+ 
    "What does that value represent in Celsius?\n"+
    "To calculate the value we need to know that C = K - 273\n"+
    "Therefore, the boiling point of water is %d Celsius\n",ebulicaoAguaC)
}
