import socket
import threading
from cryptography.fernet import Fernet
import string
import random

flag = "hz{w0w_3erNe5_1s_g00d_crypt0_@1g0r1thm}"
incorrect = "oh no, u didn't decrypt me"
banner = """

    __         ____         _ _               ____                     __ 
   / /_  ___  / / /___     (_| )____ ___     / __/__  _________  ___  / /_
  / __ \/ _ \/ / / __ \   / /|// __ `__ \   / /_/ _ \/ ___/ __ \/ _ \/ __/
 / / / /  __/ / / /_/ /  / /  / / / / / /  / __/  __/ /  / / / /  __/ /__ 
/_/ /_/\___/_/_/\____/  /_/  /_/ /_/ /_/  /_/  \___/_/  /_/ /_/\___/\__( )
                                                                       |/ 
    _ ____                            
   (_) __/  __  __   _________ _____  
  / / /_   / / / /  / ___/ __ `/ __ \ 
 / / __/  / /_/ /  / /__/ /_/ / / / / 
/_/_/     \__,_/   \___/\__,_/_/ /_( )
                                   |/ 
       __                           __                         ____ 
  ____/ /__  ____________  ______  / /_   ____ ___  ___     _ / __ \
 / __  / _ \/ ___/ ___/ / / / __ \/ __/  / __ `__ \/ _ \   (_) / / /
/ /_/ /  __/ /__/ /  / /_/ / /_/ / /_   / / / / / /  __/  _ / /_/ / 
\__,_/\___/\___/_/   \__, / .___/\__/  /_/ /_/ /_/\___/  (_)_____/  
                    /____/_/                                   

"""


chars = string.ascii_letters + string.digits + string.punctuation
random_value = ''.join([random.choice(chars) for i in range(10)])

print(random_value)

key = "PnMBraMy8Sy-3xbC7olJ7HeMgqiKgY03vGtwXcIrSzQ="
key = bytes(key,'utf-8')
f = Fernet(key)
byte_data  = bytes(random_value,'utf-8')
decrypt_me = f.encrypt(byte_data)
size = 1024
HEADER = 64
ADDR = ('0.0.0.0', 9000)
FORMAT = 'utf-8'
server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind(ADDR)



def handle_client(conn, addr):
    conn.settimeout(5)
    for i in range(10):
        try:
            data = conn.recv(size).decode()
            if random_value in data:
                conn.send(flag.encode())
            else:
                conn.send(incorrect.encode())
                conn.close()
        except:
            conn.send(incorrect.encode())
            conn.close()
    conn.close()




def start():
    server.listen()
    print("Server is listening on")

    conn, addr = server.accept()
    conn.send(banner.encode())
    conn.send(str(decrypt_me).encode())
    thread = threading.Thread(target=handle_client, args=(conn, addr))
    thread.start()



print("STARTING")
start()
