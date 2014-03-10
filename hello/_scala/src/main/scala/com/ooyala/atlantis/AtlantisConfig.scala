package com.ooyala.atlantis
import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.scala.DefaultScalaModule
import java.io.File

object AtlantisConfig {
  case class ContainerConfig(id: String, host: String, env: String)
  case class AppConfig(http_port:Int, secondary_ports:Array[Int], container:ContainerConfig, dependencies:Map[String,Map[String,Any]])
  def load() : AppConfig = {
    val mapper = new ObjectMapper()
    mapper.registerModule(DefaultScalaModule)
    return mapper.readValue(new File("/etc/atlantis/config/config.json"), classOf[AppConfig])
  }
}
