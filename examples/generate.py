import os
import json
import networkx as nx

# Implement all from https://networkx.org/documentation/stable/reference/readwrite/index.html

graphs = {
    "complete": nx.complete_graph(n=5),
    "lollipop": nx.lollipop_graph(m=5, n=3, create_using=None),
    "balanced_tree": nx.balanced_tree(r=5, h=3),
}

def out(data, path):
    with open(path, "w") as file: 
        json.dump(data, file)


def json_graph(G, path):
    data = nx.node_link_data(G)
    out(data, path)


def json_adjacency(G, path):
    data = nx.adjacency_data(G)
    out(data, path)

def json_cytoscape(G, path):
    data = nx.cytoscape_data(G)
    out(data, path)



writer = {
    "adjlist": nx.write_edgelist,
    "multiline-adjlist": nx.write_multiline_adjlist,
 #   "dot": nx.write_dot,
    "edgelist": nx.write_edgelist,
    "gexf": nx.write_gexf,
    "gml": nx.write_gml,
    "graphml": nx.write_graphml,
    "json": json_graph,
    "adjacency.json": json_adjacency,
    "cytoscape.json": json_cytoscape,
    "net": nx.write_pajek
}

# TODO IMPLEMENT MORE METHODS

for g_name, G in graphs.items():
    path = g_name
    if not os.path.exists(path):
        os.makedirs(path)

    for name, write in writer.items():
        filename = f"{g_name}/graph.{name}"
        print(G, filename)
        write(G, filename)

    