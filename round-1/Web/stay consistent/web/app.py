from flask import Flask
from uuid import uuid1

app = Flask(__name__)
flag = "hz{yeaa_finally_u_got_it}"
first_id = str(uuid1())
search_id = {}

search_id[first_id] = flag
me = str(uuid1())
return_id = me
search_id[me] = first_id

for i in range(1000):
    tmp = me
    me = str(uuid1())
    search_id[me] = tmp

@app.route('/')
def index():
    return  """
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        
        <title>hello world</title>
    </head>
    <body>
        
       <div>
            Go to /id/{}
       </div>
    </body>
    </html>
    """.format(me)

@app.route('/id/<id>')
def flag(id):
    return  """
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>hello world</title>
    </head>
    <body>
        
       <div>
            Go to /id/{}
       </div>
    </body>
    </html>
    """.format(search_id[id])
    

if __name__ == "__main__":
    app.run(host='0.0.0.0')
