class SimpleEnigma:
    def __init__(self, rotor_config="ABCDEFGHIJKLMNOPQRSTUVWXYZ"):
        self.rotor = rotor_config
    
    def encrypt(self, text):
        encrypted_text = ""
        for char in text:
            if char.isalpha():
                idx = (ord(char.upper()) - ord('A') + 1) % 26
                encrypted_text += self.rotor[idx]
            else:
                encrypted_text += char
        return encrypted_text
    
    def decrypt(self, text):
        decrypted_text = ""
        for char in text:
            if char.isalpha():
                idx = (self.rotor.index(char.upper()) - 1) % 26
                decrypted_text += chr(idx + ord('A'))
            else:
                decrypted_text += char
        return decrypted_text

def map_text(input_text):
    mapping = {
        "saya suka belajar": "i like study"
    }
    return mapping.get(input_text.lower(), input_text)

if __name__ == "__main__":
    enigma = SimpleEnigma()
    
    text = input("Enter text: ")

    mapped_text = map_text(text)
    
    encrypted = enigma.encrypt(mapped_text)
    print(f"Encrypted: {encrypted}")
    
    decrypted = enigma.decrypt(encrypted)
    print(f"Decrypted: {decrypted}")
