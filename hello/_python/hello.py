import os.path
from flask import Flask, request, make_response

app = Flask(__name__)

@app.route('/')
def hello():
  return "<!DOCTYPE html><html><head><title>hello-go</title></head><body><pre>\n"+request.path+"\n</pre></body></html>\n"

@app.route('/healthz')
def healthz():
  status = 'OK'
  code = 200
  if os.path.exists('/etc/maint'):
    status = 'MAINTENANCE'
    code = 404
  response = make_response(status, code)
  response.headers['Server-Status'] = status
  return response

if __name__ == '__main__':
  app.run(debug=True)
