from flask import Flask, request, render_template

app = Flask(__name__)

@app.route('/')
def home():
   return render_template('login.html')

@app.route('/flag', methods=['GET'])
def query_strings():

    args1 = request.args.get('user')
    args2 = request.args.get('pass')

    if((args1=="sage") and (args2=="Password123!")):
        return "hzu18{You_G0t_th3_fl@g}"
    else:
        return "Username or Password is Incorrect in here"


if __name__ == '__main__':
   app.run(host='0.0.0.0')
