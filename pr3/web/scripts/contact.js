function updFormAction() {
    document.querySelector("form").action =
        "mailto:korobkov200364@gmail.com?body="+
        encodeURIComponent(document.querySelector("textarea").value+
        "\nReply at: ")+
        encodeURIComponent(document.querySelector('input[name="email"]').value+
        "\n\nBest regards,\n")+
        encodeURIComponent(document.querySelector('input[name="name"]').value);
}
