# project-quadtree

Code source initial pour le projet d'introduction au développement (R1.01) et de SAÉ implémentation d'un besoin client (SAE1.01), année 2023-2024.

# Le Projet Etudiant

Réalisé par Andrea Bianchi et Tanguy Branellec (BUT1 TP4.2).

# Implémenté 
- Eau animée
- Eau bloquante
- Génération aléatoire (algo de prise de bloc voisin)
- Génération par biome (Perlin Simplifié)
- Enregistrement de la carte générée
- Minicarte des biomes (+ point sur la carte)
- Téléporteurs
## Installation
Installez go pour votre OS. Pour Linux :
```bash
$sudo apt-get update
$sudo apt-get install golang
```
Pour Windows, rendez vous sur [le site officiel de go](https://go.dev/).

Une fois Go installé, mettez vous dans le dossier cmd/ dans votre terminal, et effectuez :

```bash
go build
```

## Lancement et Configuration
Pour lancer le programme sous windows :
```bash
.\cmd.exe
```
Pour lancer le programme sous linux :
```shell
./cmd
```

Les **options de configuration** sont les suivantes :
```json
{
  "DebugMode": true,                          // debug mode (d)
  "NumTileX": 10,                             // tuiles sur écran en X
  "NumTileY": 10,                             // tuiles sur écran en Y
  "TileSize": 16,                             // taille des tuiles
  "MapTileSize": 1,                           // taille des tuiles de Map
  "NumCharacterAnimImages": 4
  "NumFramePerCharacterAnimImage": 5,
  "NumTileForDebug": 6,
  "CameraMode": 1,                            // 0 = camera fixe ; 1 = camera de suivi
  "FloorKind": 3,                             // 0 default 1 load 2 quadtree 3 biome 4 randomgen
  "FloorFile": "../floor-files/validation",   // emplacement de la carte

  "SaveMap": false,                           // sauvegarde map (quand possible)
  "ForceCenterPlayer": false,                 // positionnement à distance des limites
  "MapDimX": 60,                              // dimension minimap en X
  "MapDimY": 60,                              // dimension minimap en Y
  "MapX": 20,                                 // dimension map génération en X
  "MapY": 20                                  // dimension map génération en Y
}
```


