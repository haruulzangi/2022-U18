from flask import Flask, request
from random import randint
from Crypto.Hash import MD5


def hash_string(input):
    byte_input = input.encode()
    hash_object = MD5.new()
    hash_object.update(byte_input)
    return str(hash_object.hexdigest())


app = Flask(__name__)


@app.route("/")
def home():
    number1 = randint(1, 100)
    number2 = randint(1, 100)
    number3 = randint(1,100)
    number4 = randint(1,100)
    number5 = randint(1,100)
    # niilber = number1+number2
    global niilber
    niilber = number1*number2+number5-(number3+number4)

    experssion = hash_string(str(number1)) + " * " + hash_string(str(number2)) + " + " + hash_string(str(number5)) + " - " + "(" + hash_string(str(number3)) + " + "  + hash_string(str(number4)) + ")"

    return """
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="refresh" content="5" />
        <title>hello world</title>
    </head>
    <body>
        <h1>Дараах илэрхийлэлийг бодоорой. :D</h1>
        <form action="{}" method="POST">
            {}
            <input type="text" name="sum" >
            <input type="submit" >
        </form>
    </body>
    </html>
    """.format("/sum", experssion)


@app.route('/sum', methods=['POST'])
def sum():

    if request.form.get("sum") == str(niilber):
        return "hz{u_r_so_good_at_math}"
    else:
        return "<h1 style='color:red';> Дахиад оролдоол үзээрэй чадах болноо. :D</h1>"


if __name__ == "__main__":
    app.run()