//Sélection du form pour la redirection
const form = document.getElementById('redirect_categorie') 

// Redirection vers le groupe cliqué
function displayCategory(ID){
    // Ajout d'une action au form pour la redirection vers l'ID du groupe
    form.setAttribute('action', `/oneCategory=${ID}`)
    // Submit du form
    form.submit()
}