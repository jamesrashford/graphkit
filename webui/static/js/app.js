


$.getJSON("/static/graph.json", function(data) {
    console.log(data);

    const graph = new graphology.Graph();
    graph.import(data);
    /*
graph.addNode("1", { label: "Node 1", x: 0, y: 0, size: 10, color: "blue" });
graph.addNode("2", { label: "Node 2", x: 1, y: 1, size: 20, color: "red" });
graph.addEdge("1", "2", { size: 5, color: "purple" });
*/

    console.log(graph);
    const sigmaInstance = new Sigma(graph, document.getElementById("sigma-container"));
});

