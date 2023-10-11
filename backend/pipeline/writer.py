import json

def GovPageFile(data: dict):
    path = f'database/public/{data["title"]}.json'
    with open(path,"w") as f:
        json.dump(data,f, indent=4)
        print(f"data saved to: {path}")


