import socket

if __name__ == '__main__':
	s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
	s.connect(('localhost', 8080))

	data = s.recv(1024).decode('utf-8')
	print(data)
	data = s.recv(1024).decode('utf-8')
	print(data)

	numbers = data.split(' ')[1]
	n = int(numbers[numbers.index('-') + 1:])

	l = 1
	r = n
	while l <= r:
		mid = (l + r) // 2
		print(mid)
		s.send(f'{mid}\n'.encode('utf-8'))

		data = s.recv(1024).decode('utf-8')

		print(data)
		if 'Арай' in data:
			l = mid + 1
		elif 'Хэтрээд' in data:
			r = mid - 1
		else: break

		if not 'Хайрцагны' in data:
			data = s.recv(1024).decode('utf-8')
		print(data)

	s.close()
