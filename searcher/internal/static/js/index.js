function send() {
    let query = String(document.getElementById('query').value).trim();

    if (query !== "") {
        console.log(`/search/${query.split(" ").join("_")}`)
        location.replace(`/search/${query.split(" ").join("_")}`)
    }
}