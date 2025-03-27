package main

import (
	"client/globals"
	"client/utils"
	"log"
 	"os"
	"bufio" 
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Soy un Log" usando la biblioteca log
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	log.Println("Soy un Log")

	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil { //TODO cuando ocurre?
		log.Fatalf("No se pudo cargar la configuración")
	}

	// loggeamos el valor de la config
	log.Println("Mensaje desde config.json:", globals.ClientConfig.Mensaje)
	log.Println("ip desde config.json:", globals.ClientConfig.Ip)
	log.Println("puerto desde config.json:", globals.ClientConfig.Puerto)

	//lea de consola todas las líneas que se ingresen, las loguee y, si se ingresa una línea vacía, termine con el programa.
	reader := bufio.NewReader(os.Stdin) // no uso utils.LeerConsola() porque necesito text
	text, _ := reader.ReadString('\n')
	log.Println(text)
	
	for text != "\n" && text != "\r\n" {
		text, _ = reader.ReadString('\n')
		log.Println(text)
	}
	
	log.Println("Salida del programa.")


	/* utils.LeerConsola() */

	/* 
	func LeerConsola() {
	reader := bufio.NewReader(os.Stdin)
	log.Println("Ingrese los mensajes")
	text, _ := reader.ReadString('\n')
	log.Print(text)
}
	*/


/* 	utils.LeerConsola() */

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	/* // enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	// leer de la consola el mensaje
	utils.LeerConsola()
	log.Println("Enviando mensaje al servidor...")

	
	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete()
 */


}