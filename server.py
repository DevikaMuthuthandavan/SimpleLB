from flask import Flask
app = Flask(__name__)

serverName = 'Server 1'

@app.route('/')
def hello_world():
   return serverName

if __name__ == '__main__':
   app.run()
