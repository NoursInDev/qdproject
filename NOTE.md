# Notes à l'attention d'Andrea:

## Methodes de création des terrains
1. Via Floor, prenant légèrement plus de place, toute première partie du projet.
2. Via QuadTree, plus optimisé en terme de place.

## Utilisation des QuadTrees
-> Les QuadTrees sont définis par 4 branches qui elles-mêmes donnent 4 branches et ainsi de suite jusqu'à ce que la branche n'ai qu'un seul type de blocs.

-> Le QuadTree est défini par q. Chaque branche est définit comme un noeud (node) et appelée via un pointeur *node.

    -> Pour appeler la première branche (la branche initiale), il faut faire q.root, qui renvoies à la node initiale.

## Fonctionnement Général
Le fichier init.go lance character/init.go -> init ; camera/init.go -> init ; floor/init.go -> init

Le fichier floor/init.go permet d'appeler au choix du config.json soit les fichiers de génération QuadTree soit ceux de génération Floor classique.



# Suite des opérations:

-> Problèmes et modifs actuellement à faire :
    
    -> Créer un fichier propre à la sauvegarde de la carte + intégrer des fonctions d'enregistrement dynamique (avec 6.12)

    -> Patch le problème de sortie de map en mode QuadTree (je m'en occuperai tkt)

-> Faire les extensions:

    -> 6.1 créer un nouveau mode de map : génération aléatoire d'un terrain (de taille donnée) - FAIT -> EXTENSION=> 6.2 enregistrement de la map - FAIT -> EXTENSION=> 6.12 Generation de terrain progressive -> EXTENSION=> (aucun) Génération par biome / modes de génération - FAIT|| uniquement un changement de génération aléatoire, interressant mais pas de foufou non plus

    -> 6.3 animation des décors - FAIT (EAU ANIMEE)

    -> 6.5 interdiction de marcher sur l'eau - FAIT -> EXTENSION=> (aucun) Mettre une animation spécifique de bateau / ajouter une mécanique de saut => PLOUF ET RESPAWN EN 0-0 || légèrement valorisé

    -> 6.7 fluidification de la caméra

    -> 6.9 affichage du terrain amélioré

-> Faire les tests

-> Faire les extensions personnelles:

    -> mode multijoueur -> EXTENSION=> chat vocal de proximité || Pas rentable en temps, à faire en dernier

    -> menu au démarrage || à voir mais relativement interressant

    -> ajouter un Easter Egg > Konami Code => Lance une animation de cuisine (LET HIM COOK NOW)

    -> waypoints (!= tp)

    -> minimap des chunks générés || Interressant - FAIT