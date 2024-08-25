import os
import networkx as nx

# Implement all from https://networkx.org/documentation/stable/reference/readwrite/index.html

graphs = {
    "complete": nx.complete_graph(n=5),
    "lollipop": nx.lollipop_graph(m=5, n=3, create_using=None),
    "balanced_tree": nx.balanced_tree(r=5, h=3),
}

writer = {
    "adjlist": nx.write_edgelist,
    "multiline-adjlist": nx.write_multiline_adjlist,
    "dot": nx.write_dot,
    "edgelist": nx.write_edgelist,
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

    