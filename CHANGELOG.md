***25/10/23***
> Dev Fonction floor/init.go -> readFloorFromFile
> 
> Dev Fonction floor/update.go -> updateFromFileFloor

***26/10/23***
> Dev Fonction quadtree/make.go -> MakeFromArray & BuildQuadTree
> 
> Dev Fonction quadtree/get.go -> GetContent & ReturnFloor
> 
> Création des fichiers NOTE.md et CHANGELOG.md
> 
> Mise à jour du README.md

***27/10/23***
> Définition NOTE.md -> Liste des Objectifs (suite des opérations)
> 
> Dev Fonction floor/init.go -> floorGeneration
> 
> Ajout floor/floor.go -> constante randomFloor
> 
> Dev Fonction floor/init.go -> SaveFloorToFile

***28/10/23***
> Ajout option cmd/config.json -> SaveMap

***13/11/23***
> Setup Animation Eau (simultané)
> 
> Setup Animation Eau (asynchrone) -> UPDATE PREVIOUS

***22/11/23***
> Fix QuadTrees impairs
> 
> Fix QuadTrees surgénération (hors limite)
> 
> Blocage de l'eau
> 
> Finalisation Quadtrees

***23/11/23***
> Ajout package noise -> Bruit de Perlin (simplifié)
> 
> Création fonction Perlin2DMap
> 
> Création Génération via Biomes (profil de génération 3)

***28/12/23 et avant***
> Ajout package minimap (affichage de minimap de biome)
> 
> Création de map.go, draw.go et mapids.png
> 
> Modification de Draw() dans game.go pour l'affichage
> 
> Création fonction Draw() pour la minimap
> 
> Ajout des dimensions de tuile de map et de la taille de la minimap
> 
> Positionnement des limites de la minimap

***09/01/24***
> Création de Release
> 
> Patch de bugs minimap

***10/01/24***
> Création de Téléporteurs
> 
> Création d'animations pour les téléporteurs
> 
> Modification de blocking.go
> 
> Création de Release