import os
import re
import requests
import ftfy
import pandas as pd
from bs4 import BeautifulSoup

def parseTitle(text):
    title_parts = text.split('-')
    title = None
    if len(title_parts) == 3:
        matches = re.findall(r'\"(.+?)\"', title_parts[1].strip())
        if len(matches) == 1:
            title = matches[0]
    return title

def parse(content, encoding, season, episode):
    soup = BeautifulSoup(content, 'html.parser', from_encoding=encoding)
    title = parseTitle(soup.title.text)
    if title is None:
        if season == 1 and episode == 1:
            title = "Pilot"
    quotes = soup.find_all('div','quote')
    df = pd.DataFrame(columns = ['season','episode','episode_name', 'scene', 'character','quote'])
    for i, quote_div in enumerate(quotes, start=1):
        characterList = []
        quoteList = []
        character = None
        quote = None
        for child in quote_div.children:
            if child.name == 'b' and child.string is not None:
                character = child.string[:-1]
            else:
                if child.string is not None:
                    parsed_element = child.string.strip()
                    quote = parsed_element if len(parsed_element) > 0 else None

            if character is not None and quote is not None:
                characterList.append(character)
                quoteList.append(ftfy.fix_encoding(quote))
                character = None
                quote = None

        if len(characterList) == len(quoteList):
            seasonList = [season]*len(characterList)
            episodeList = [episode]*len(characterList)
            episodeNameList = [title]*len(characterList)
            sceneList = [i]*len(characterList)
            temp_df = pd.DataFrame(
                {
                    'season': seasonList, 
                    'episode': episodeList, 
                    'episode_name': episodeNameList,
                    'scene': sceneList,
                    'character': characterList,
                    'quote': quoteList
                }
            )
            df = df.append(temp_df)

    return df

if __name__ == "__main__":
    episodes = [6, 22, 23, 14, 26, 24, 24, 24, 23]
    domain = 'https://www.officequotes.net'
    for i, j in zip(range(1,10), episodes):
        for k in range(1, j+1):
            episode = str(k) if k >= 10 else '0{}'.format(k)
            pagename = '{}/no{}-{}.php'.format(domain, i, episode)
            r = requests.get(pagename)
            encoding = r.encoding if "charset" in r.headers.get("content-type", "").lower() else None
            df = parse(r.content, encoding, i, int(episode))
            json_filename = 'season_{}_episode_{}.json'.format(i, episode)
            df.to_json(os.path.join(os.getcwd(), 'quotes', json_filename), orient='records', lines=True)
            