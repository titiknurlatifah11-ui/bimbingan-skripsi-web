#Oserror
import os
try:
    file = open("abcd.txt", "r")
    file.close()
except OSError as error:
    print(error)
    print("File descriptor is not associated with any terminal device")

# NameError
def message():
	try:
		name = "hi helo"
		return HiHello
	except NameError:
		return "NameError occurred. Some variable isn't defined."

print(message())

#KeyError
subjects = {'Titik': 'MTK', 
			'Eha': 'Biologi', 
			'Nathan': 'Fisika', 
			'Kekez': 'B.indo'} 
try: 
	print('subject of Ratu is:', 
		subjects['Ratu']) 
except KeyError: 
	print("Ratu's records doesn't exist") 

#TypeError
list_nama = ("titik", "nathan")
indices = [0, "1"]
for i in range(len(indices)):
    try:
        print(list_nama[indices[i]])
    except TypeError:
        print("TypeError: Check list of indices")

#EOFError
try: 
	n = int(input()) 
	print(n * 10) 
	
except EOFError as e: 
	print(e)