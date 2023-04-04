import requests

API_BASE_URL = "https://pokeapi.co/api/v2/pokemon/"

def get_json(url):
    r = requests.get(url)
    if r.status_code != requests.codes.ok:
        return None
    json_data = r.json()
    return json_data

def get_abilities(name):
    name = name.lower()
    url = API_BASE_URL + name
    json_data = get_json(url)
    if not json_data:
        return []
    abilities_list = []
    for abilities in json_data['abilities']:
        for key, value in abilities['ability'].items():
            if key == 'name':
                abilities_list.append(value)
    return abilities_list

def print_abilities(pokemon_name, abilities):
    if not abilities:
        print("no such pokemon..")
        return
    print("Name:", pokemon_name)
    print("Abilities:")
    for ab in abilities:
        print("-", ab.capitalize())

def main():
    pokemon_name = input("Enter the name of a Pokemon: ")
    abilities = get_abilities(pokemon_name)
    print_abilities(pokemon_name, abilities)

if __name__ == '__main__':
    main()
