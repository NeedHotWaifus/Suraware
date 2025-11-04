import random
import os
import sys
from pathlib import Path

def encrypt_file(input_file, output_file):
    """Encrypt a Python file using XOR cipher with random key"""
    
    # Read the input file
    with open(input_file, 'r', encoding='utf-8') as f:
        contents = f.read()
    
    # Generate random 100-digit key
    xnd = ""
    for _ in range(100):
        xnd += str(random.randint(0, 9))
    
    # XOR encrypt the contents
    no_of_itr = len(contents)
    output_string = ""
    for i in range(no_of_itr):
        current_string = contents[i]
        current_key = xnd[i % len(xnd)]
        output_string += chr(ord(current_string) ^ ord(current_key))
    
    # Create the encrypted stub
    encrypted_repr = repr(output_string).replace("'", "")
    
    stub_code = f"""wopvEaTEcopFEavc = \"{encrypted_repr}\"

iOpvEoeaaeavocp = \"{xnd}\"

uocpEAtacovpe = len(wopvEaTEcopFEavc)
oIoeaTEAcvpae = ""
for fapcEaocva in range(uocpEAtacovpe):
    nOpcvaEaopcTEapcoTEac = wopvEaTEcopFEavc[fapcEaocva]
    qQoeapvTeaocpOcivNva = iOpvEoeaaeavocp[fapcEaocva % len(iOpvEoeaaeavocp)]
    oIoeaTEAcvpae += chr(ord(nOpcvaEaopcTEapcoTEac) ^ ord(qQoeapvTeaocpOcivNva))

eval(compile(oIoeaTEAcvpae, '<string>', 'exec'))
"""
    
    # Write encrypted stub
    with open(output_file, 'w', encoding='utf-8') as f:
        f.write(stub_code)
    
    return True

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python crypter.py <input_file> [output_file]")
        sys.exit(1)
    
    input_file = sys.argv[1]
    output_file = sys.argv[2] if len(sys.argv) > 2 else "stub.py"
    
    if not os.path.exists(input_file):
        print(f"Error: Input file '{input_file}' not found")
        sys.exit(1)
    
    print(f"[*] Encrypting {input_file}...")
    print(f"[*] Generating random encryption key...")
    
    if encrypt_file(input_file, output_file):
        print(f"[+] File successfully encrypted to {output_file}")
    else:
        print(f"[-] Encryption failed")
        sys.exit(1)
