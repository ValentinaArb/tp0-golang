package main

import (
	"client/globals"
	"client/utils"
	"log"
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
	

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	// leer de la consola el mensaje
	/* utils.LeerConsola() */
	log.Println("Enviando mensaje al servidor...")

	paquete := utils.LeerConsola()
	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete(paquete, globals.ClientConfig.Ip,globals.ClientConfig.Puerto)
}