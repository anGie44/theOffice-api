#!/bin/bash

baseURL=https://www.officequotes.net/
n=(6 22 23 14 26 24 24 24 23)
for i in {1..9}; do 
   idx=$((i - 1))
   n_episodes=${n[$idx]}
   for j in $(seq 1 $n_episodes); do
   	episode=0
	if [ "$j" -lt 10 ]; then
		episode+=$j
        else 
		episode=$j
	fi
	url=${baseURL}no${i}-${episode}.php
   	curl ${url} -o season_${i}_episode_${episode}.php
   done
done

