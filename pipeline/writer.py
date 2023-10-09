import json

def GovPageFile(data: dict, path):
    
    with open(path,"w") as f:
        json.dump(data,f, indent=4)
        print(f"data saved to: {path}")


