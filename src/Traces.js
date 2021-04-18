function createTracesFromJSON(ammoData) {
    let traces = []

    const baseTrace = {
        type: 'scatter3d',
        mode: 'lines+markers+text',
        textposition: 'top center',
        marker: {size: 3},
        hovertemplate:
            '<b><i>%{text}</i></b><br>' +
            'Damage: %{x}<br>' +
            'Pen: %{y}<br>' +
            'Cost: ₽ %{z}<br>'
    }

    for (const [caliber, _] of Object.entries((ammoData.record))) {
        let trace =  {
            ...baseTrace,
            name: caliber,
            x:[],
            y:[],
            z:[],
            text: []
        }

        for (const [_, ammo] of Object.entries(ammoData.record[caliber])) {
            trace.x.push(ammo.damage)
            trace.y.push(ammo.penetration)
            if (ammo.price > 5000) {
                trace.z.push(5000)
            } else {
                trace.z.push(ammo.price)
            }
            trace.text.push(ammo.name)
        }

        console.log(trace)
        traces.push(trace)
    }

    return traces
}

export default createTracesFromJSON
