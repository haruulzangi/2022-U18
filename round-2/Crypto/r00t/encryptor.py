from Crypto.Util.number import *

e = 3

def read_flag():
    with open('flag.txt', 'rb') as f:
        return f.read().strip()
    f.close()

def generate_primes():
    return getPrime(1024), getPrime(1024)

def encrypt():
    p, q = generate_primes()
    n = (p * q)

    flag = read_flag() + b"\x00" * 200
    return pow(bytes_to_long(flag),e,n), n

if __name__ == '__main__':
    encrypted_flag, n = encrypt()
    print('Encrypted flag: {}\nN: {}\n'.format(encrypted_flag, n))
