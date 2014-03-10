package com.ooyala.hello
import org.eclipse.jetty.server.Server
import org.eclipse.jetty.servlet.{DefaultServlet, ServletContextHandler}
import org.eclipse.jetty.webapp.WebAppContext
import org.scalatra.servlet.ScalatraListener
import com.ooyala.atlantis.AtlantisConfig

object JettyLauncher { // this is my entry object as specified in sbt project definition
  def main(args: Array[String]) {
    var port = 9876
    try {
      val config = AtlantisConfig.load()
      port = config.http_port
    } catch {
      case ex: Exception => {
        // do nothing
        println(ex)
      }
    }

    val server = new Server(port)
    val context = new WebAppContext()
    context setContextPath "/"
    context.setResourceBase("src/main/webapp")
    context.addEventListener(new ScalatraListener)
    context.addServlet(classOf[DefaultServlet], "/")

    server.setHandler(context)

    server.start
    server.join
  }
}
