const svg = d3.select("body").append("svg")
            .attr("width", window.innerWidth)
            .attr("height", window.innerHeight);

const tooltip = d3.select("#tooltip");

window.addEventListener("resize", () => {
    svg.attr("width", window.innerWidth).attr("height", window.innerHeight);
});

const zoomLayer = svg.append("g");

svg.call(d3.zoom()
    .scaleExtent([0.5, 5])
    .on("zoom", (event) => {
        zoomLayer.attr("transform", event.transform);
    }));

const endpoint = "/graph.json"; // Replace with your endpoint URL
d3.json(endpoint).then(data => {
    const { nodes, links, directed } = data;

    const degreeMap = {};
    links.forEach(link => {
        degreeMap[link.source] = (degreeMap[link.source] || 0) + 1;
        degreeMap[link.target] = (degreeMap[link.target] || 0) + 1;
    });
    nodes.forEach(node => {
        node.degree = degreeMap[node.id] || 0;
    });

    if (directed) {
        svg.append("defs").append("marker")
            .attr("id", "arrowhead")
            .attr("viewBox", "0 -5 10 10")
            .attr("refX", 10) // Reference point for the arrow
            .attr("refY", 0)
            .attr("markerWidth", 6)
            .attr("markerHeight", 6)
            .attr("orient", "auto")
            .append("path")
            .attr("d", "M0,-5L10,0L0,5")
            .attr("class", "arrowhead");
    }

    const simulation = d3.forceSimulation(nodes)
        .force("link", d3.forceLink(links).id(d => d.id).distance(100))
        .force("charge", d3.forceManyBody().strength(-50))
        .force("center", d3.forceCenter(window.innerWidth / 2, window.innerHeight / 2))
        .on("tick", ticked);

    const link = zoomLayer.append("g")
        .attr("class", "links")
        .selectAll("line")
        .data(links)
        .enter().append("line")
        .attr("stroke", "#999")
        .attr("stroke-opacity", 0.6)
        .attr("marker-end", directed ? "url(#arrowhead)" : null);

    const node = zoomLayer.append("g")
        .attr("class", "nodes")
        .selectAll("circle")
        .data(nodes)
        .enter().append("circle")
        .attr("r", d => 5 + d.degree)
        .attr("fill", "#69b3a2")
        .on("mouseover", (event, d) => showTooltip(event, d))
        .on("mousemove", (event) => moveTooltip(event))
        .on("mouseout", hideTooltip)
        .call(drag(simulation));

    node.append("title")
        .text(d => d.id);

    function ticked() {
        link.attr("x1", d => d.source.x)
            .attr("y1", d => d.source.y)
            .attr("x2", d => calculateArrowheadX(d))
            .attr("y2", d => calculateArrowheadY(d));

        node.attr("cx", d => d.x)
            .attr("cy", d => d.y);
    }

    function calculateArrowheadX(d) {
        const dx = d.target.x - d.source.x;
        const dy = d.target.y - d.source.y;
        const distance = Math.sqrt(dx * dx + dy * dy);
        const radius = 5 + d.target.degree;
        return d.target.x - (dx / distance) * radius;
    }

    function calculateArrowheadY(d) {
        const dx = d.target.x - d.source.x;
        const dy = d.target.y - d.source.y;
        const distance = Math.sqrt(dx * dx + dy * dy);
        const radius = 5 + d.target.degree;
        return d.target.y - (dy / distance) * radius;
    }

    function drag(simulation) {
        return d3.drag()
            .on("start", (event, d) => {
                if (!event.active) simulation.alphaTarget(0.3).restart();
                d.fx = d.x;
                d.fy = d.y;
            })
            .on("drag", (event, d) => {
                d.fx = event.x;
                d.fy = event.y;
            })
            .on("end", (event, d) => {
                if (!event.active) simulation.alphaTarget(0);
                d.fx = null;
                d.fy = null;
            });
    }

    function showTooltip(event, d) {
        tooltip
            .style("opacity", 1)
            .html(`ID: ${d.id}<br>Degree: ${d.degree}`);
    }

    function moveTooltip(event) {
        tooltip
            .style("left", (event.pageX + 10) + "px")
            .style("top", (event.pageY + 10) + "px");
    }

    function hideTooltip() {
        tooltip
            .style("opacity", 0);
    }
}).catch(error => {
    console.error("Error fetching or parsing the JSON data:", error);
});