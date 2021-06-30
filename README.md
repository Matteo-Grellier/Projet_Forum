# 📑 Projet_Forum

Ce projet consiste en la création d'un forum en ligne qui reprend les caractéristiques d'un forum classique : la publication de posts par des utilisateurs, like ou dislike de posts, pouvoir s'inscrire/se connecter, une multitude de catégories, la création de topics etc...

# 📝 Consignes

L'utilisateur non connecté pourra : 

- Lire des sujets, posts, commentaires

L'utilisateur connecté devra avoir la possibilité de :

- Créer des sujets (liés à une catégorie)

- Réagir aux posts du sujet (likes, dislikes, commentaires)

Un système de filtrage des sujets est mis en place :

- Par catégorie

- Les sujets que l'utilisateur aura liké ou posté



## Fonctionnalités attendues

- L'utilisateur peut s'inscrire et se connecter

- Les mots de passe seront *hashés*

- Une session utilisateur sera mise en place en utilisant un cookie avec un temps d'expiration

## Contraintes 

- Serveur web en **Golang**

- Une **URL** par page

- Base de donnée gérée & administrée avec **SQLite**

- Les packages autorisés :
     - Packages standards
     - **bcrypt** (mots de passe sécurisés)
     - **sqlite3**
     - **uuid** (sessions de connexion avec cookies)

# 👉 Pour commencer

👉 Télécharger ``Golang``: https://golang.org 

## ☝️ Pré-requis

Avant de commencer, nous avons besoin d'installer des packages sur notre 💻 terminal en utilisant la commande suivante ``go get`` suivie de :

- ``golang.org/x/crypto/bcrypt``

- ``github.com/satori/go.uuid``

- ``github.com/mattn/go-sqlite3``


## 📥 Téléchargement 

Téléchargeons le projet : 

- Version release 👉 [ici](https://github.com/Matteo-Grellier/Projet_Forum/archive/refs/heads/main.zip)

# 🟢 Lancement 

Une fois le projet téléchargé, lançons ``Visual Studio Code``. 

Pour le lancement du serveur, mettons-nous à la racine du projet pour pouvoir exécuter le fichier suivant ``server.go``.

Dans le terminal :

🔹 ``go run server.go``

Le bon lancement du serveur se traduira par le message suivant : 

![img](https://raw.githubusercontent.com/Matteo-Grellier/Projet_Forum/README/static/images/Start.png)

Ouvrez votre navigateur et rendez-vous sur :

🔸 ``http://localhost:8080`` ou directement [ici](http://localhost:8080)

---

👉 ⚠️ *Les exécutions de pages ou les erreurs sont répertoriées dans notre 💻 terminal.*

![img](https://raw.githubusercontent.com/Matteo-Grellier/Projet_Forum/README/static/images/Error.png)

# 🔍 Architecture 

Le dossier ``Projet_Forum`` se découpe en plusieurs sous dossiers : 

![img](https://raw.githubusercontent.com/Matteo-Grellier/Projet_Forum/README/static/images/Architecture.png)

* BDD : Dossier qui regroupe l'ensemble des fonctions pour les fonctionnalités liées à la base de données. (Ajout d'un post, ajout d'un utilisateur, etc.)

* Handlers : Dossier qui regroupe l'ensemble des fonctions pour l'affichage des pages du site et les fonctionnalités plus globales. (Création du cookie, vérification des entrées lors de l'inscription, etc.)

* Templates : Dossier qui regroupe les pages HTML (Hypertext Markup Language).
    * Layouts : Dossier qui regroupe des templates utilisées dans plusieurs pages HTML. (header.html, sidebar.html)

* Static : 
    * CSS : Feuilles de style en cascade.
    * JavaScript : Dossier qui regroupe les scripts.
    * Images : Dossier qui regroupe les images.


# 🎥 Démonstration

- 👉 [Démo](https://www.youtube.com/watch?v=JvJw3lWWQ_k)

# 🖥 Réalisation

Front-End :

- <img alt="HTML5" src="https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white"/> 

- <img alt="CSS3" src="https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white"/> 

- <img alt="JavaScript" src="https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E"/>

Back-end :

- <img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white"/>

Base de données :

- <img alt="SQLite" src ="https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white"/>

<!-- Conteneur : 

- <img alt="Docker" src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white"/> -->


# ⚙️ Version

Liste des versions :

[![Generic badge](https://img.shields.io/static/v1?label=DERNIERE&message=VERSION&color=<green>?style=flat-square)](https://shields.io/)
- 👉 [v4.0](https://github.com/Matteo-Grellier/Projet_Forum/releases/tag/v4.0)

[![Generic badge](https://img.shields.io/badge/PRECEDENTE-VERSION-red)](https://shields.io/)
- 👉 [v3.0](https://github.com/Matteo-Grellier/Projet_Forum/releases/tag/v3.0)
- 👉 [v2.0](https://github.com/Matteo-Grellier/Projet_Forum/releases/tag/v2.0)
- 👉 [v1.0](https://github.com/Matteo-Grellier/Projet_Forum/releases/tag/v1.0)


# 👥 Équipe

Projet réalisé à Nantes Ynov Campus par les apprenants de la promo B1 Informatique 2020/2021

- ``Elouan DUMONT`` alias [@ByMSRT](https://github.com/ByMSRT)

- ``Mattéo GRELLIER`` alias [@Matteo-Grellier](https://github.com/Matteo-Grellier)

- ``Malo LOYER-VIAUD`` alias [@Karrwolf](https://github.com/Karrwolf) (ATTENTION Problème de GIT, Malo est aussi nommé [@LemonIceStuff](https://github.com/LemonIceStuff))

- ``Olivia MOREAU`` alias [@Liv44](https://github.com/Liv44)

Lien des contributions 👉 [ici](https://github.com/Matteo-Grellier/Projet_Forum/graphs/contributors).

***
*Nantes YNOV Campus - B1 Informatique - 2020/2021*