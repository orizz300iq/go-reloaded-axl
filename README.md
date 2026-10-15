# Go-Reloaded-axl

## Description
**Go-Reloaded** est un outil simple de complétion, d'édition et d'autocorrection de texte développé en Go. Il permet de traiter automatiquement un fichier texte en appliquant un ensemble de transformations syntaxiques, typographiques et grammaticales (conversions de bases, modifications de casse, gestion de la ponctuation, des guillemets et des articles)[cite: 1, 2].

---

## Fonctionnalités

L'outil gère les modifications suivantes :
* **Conversions numériques** : 
  * `(hex)` : Convertit le nombre hexadécimal précédent en sa version décimale[cite: 1].
  * `(bin)` : Convertit le nombre binaire précédent en sa version décimale[cite: 1].
* **Modifications de casse** :
  * `(up)` : Met le mot précédent en majuscules[cite: 1].
  * `(low)` : Met le mot précédent en minuscules[cite: 1].
  * `(cap)` : Met une majuscule initiale au mot précédent[cite: 1].
  * `(tag, <number>)` : Applique la modification (`up`, `low` ou `cap`) sur les $N$ mots précédents (ex: `(up, 2)`)[cite: 1, 2].
* **Ponctuation et typographie** :
  * Colle les signes de ponctuation (`.`, `,`, `!`, `?`, `:`, `;`) au mot précédent et ajoute un espace après[cite: 2].
  * Gère et normalise les groupes de ponctuation (ex: `...`, `!?`, `?!`)[cite: 2].
  * Positionne correctement les guillemets simples (`'`) de part et d'autre des mots concernés[cite: 2].
* **Grammaire** :
  * Transforme automatiquement l'article `a` (ou `A`) en `an` (ou `An`) lorsque le mot suivant commence par une voyelle (`a, e, i, o, u`) ou par un `h`[cite: 2].

---

## Prérequis

* Langage **Go** installé sur votre machine (version 1.18 ou supérieure recommandée).

---

## Utilisation

Le programme prend en argument le fichier d'entrée contenant le texte à modifier et le fichier de sortie où sera écrit le résultat[cite: 1, 2].

Exécutez la commande suivante dans votre terminal :

```bash
go run . <fichier_depart> <nom_fichier_arrivee>