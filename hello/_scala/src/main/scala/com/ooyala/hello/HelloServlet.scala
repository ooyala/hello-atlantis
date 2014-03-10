package com.ooyala.hello

import org.scalatra._
import scalate.ScalateSupport

class HelloServlet extends HelloscalaStack {

  get("/*") {
    <html>
      <body>
        <pre>Hello from Scala</pre>
      </body>
    </html>
  }

}
