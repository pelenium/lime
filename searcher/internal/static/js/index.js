function send() {
    var query = String(document.getElementById('query').value);

    if (query !== "") {
        console.log(query);

        var jsn = {
            query: query,
        };

        fetch("/api/:req", {
            method: "POST",
            body: JSON.stringify(jsn),
            headers: {
                "Content-Type": "application/json"
            }
        })
            .then(response => response.json())
            .then(data => {
                console.log("Ответ от сервера:", data);
            })
            .catch(error => {
                console.error("Ошибка при отправке данных:", error);
            });

        document.getElementById('request').value = "";
    }
}