def vigenere_table():
    table = []
    for i in range(26):
        table.append([])
    for row in range(26):
        for column in range(26):
            if (row + 65) + column > 90:
                table[row].append(chr((row+65) + column -26))
            else:
                table[row].append(chr((row + 65) + column))
    return table

def vigenere_userInput_and_userKey():
    userInput = input('Enter PlainText: ').upper()
    userKey = 'HARUULZANGI'
    addKey = 7
    updatedKey = ''
    k = 0
    for i in range(len(userInput)):
        if ord(userInput[i]) == 32:
            updatedKey = updatedKey + ' '
        else:
            if k < len(userKey):
                updatedKey = updatedKey + userKey[k]
                k = k + 1
            else:
                k = 0
                updatedKey = updatedKey + userKey[k]
                k = k + 1
    return userInput, updatedKey, addKey

def vigenere_encryption(userInput, updatedKey):
    table = vigenere_table()
    vigEncryption = ''
    for i in range(len(userInput)):
        if userInput[i] == chr(32):
            vigEncryption = vigEncryption + ' '
        else:
            row = ord(userInput[i]) - 65
            column = ord(updatedKey[i]) - 65
            vigEncryption = vigEncryption + table[row][column]
            vigEncryptedText = '{}'.format(vigEncryption)
    return vigEncryptedText

def ceaser_key_diffusion(updatedKey, addKey):
    updatedKey2 = ''
    for i in range(len(updatedKey)):
        updatedKey2 = updatedKey2 + chr(((ord(updatedKey[i]) + addKey) - 65) % 26 + 65)
    return updatedKey2

def ceaser_cipher_enc(text, addKey):
    encryptedText = ''
    for i in range(len(text)):
        if ord(text[i]) == 32:
            encryptedText = encryptedText + ' '
        else:
            encryptedText = encryptedText + chr(((ord(text[i]) + addKey) - 65) % 26 + 65)
    return encryptedText

def itr_count(updatedKey, userInput):
    counter = 0
    result = ''
    for i in range(26):
        if updatedKey + i > 90:
            result = result + chr(updatedKey + (i-26))
        else:
            result = result + chr(updatedKey + i)
    for i in range(len(result)):
        if result[i] == chr(userInput):
            break
        else:
            counter = counter + 1

    return counter

def main():
    userInput, updatedKey, addKey = vigenere_userInput_and_userKey()
    encryption1 = vigenere_encryption(userInput, updatedKey)
    keyDiff = ceaser_key_diffusion(updatedKey, addKey)
    encryption2 = vigenere_encryption(encryption1, keyDiff)
    encryption3 = ceaser_cipher_enc(encryption2, addKey)
    print('                        ')
    print('################################################################################')
    print('Encrypted Text: ', encryption3)

if __name__ == "__main__":
    main()
