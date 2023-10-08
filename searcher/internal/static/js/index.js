function send() {
    let query = String(document.getElementById('query').value).trim();

    if (query !== "") {
        console.log(`/search/${query.split(" ").join("_")}`)
        location.replace(`/search/?req=${query.split(" ").join("_")}`)
    }
}