function createTracesFromJSON(ammoData) {
  let traces = [];

  const baseTrace = {
    type: "scatter3d",
    mode: "lines+markers+text",
    textposition: "top center",
    marker: { size: 3 },
    hovertemplate:
      "<extra></extra>" +
      "<b><i>%{text}</i></b><br>" +
      "Damage: %{x}<br>" +
      "Pen: %{y}<br>" +
      "Cost: ₽ %{z}<br>",
  };

  ammoData = ammoData["data"];

  for (const caliber in ammoData) {
    let trace = {
      ...baseTrace,
      name: caliber,
      x: [],
      y: [],
      z: [],
      text: [],
    };

    let ammoArray = [];

    for (const ammoID in ammoData[caliber]) {
      for (const _ in ammoData[caliber][ammoID]) {
        ammoArray.push(ammoData[caliber][ammoID]);
      }
    }

    let ammoName = ammoArray[0].name;
    trace.name = ammoName.substr(0, ammoName.indexOf(" ")).replace("mm", "");
    ammoArray.sort((a, b) => {
      if (a.penetration > b.penetration) return 1;
      if (a.penetration < b.penetration) return -1;
      return 0;
    });

    for (const ammo of ammoArray) {
      if (ammo.damage > 200) {
        trace.x.push(200);
      } else {
        trace.x.push(ammo.damage);
      }

      trace.y.push(ammo.penetration);

      if (ammo.price > 5000) {
        trace.z.push(5000);
      } else {
        trace.z.push(ammo.price);
      }
      trace.text.push(ammo.shortname);
    }

    if (
      caliber === "Caliber12g" ||
      caliber === "Caliber9x18PM" ||
      caliber === "Caliber762x25TT" ||
      caliber === "Caliber9x21" ||
      caliber === "Caliber30x29" ||
      caliber === "Caliber366TKM" ||
      caliber === "Caliber762x35" ||
      caliber === "Caliber1143x23ACP" ||
      caliber === "Caliber127x108" ||
      caliber === "Caliber23x75" ||
      caliber === "Caliber127x55" ||
      caliber === "Caliber545x39" ||
      caliber === "Caliber57x28" ||
      caliber === "Caliber762x54R" ||
      caliber === "Caliber86x70" ||
      caliber === "Caliber40x46" ||
      caliber === "Caliber20g"
    ) {
      trace.visible = "legendonly";
    }

    traces.push(trace);
  }

  return traces;
}

export default createTracesFromJSON;
