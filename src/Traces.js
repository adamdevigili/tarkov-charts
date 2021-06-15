export function createTraces3D(ammoData) {
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

        let ammoArray = []

        for (const [_, ammo] of Object.entries(ammoData.record[caliber])) {
            ammoArray.push(ammo)
        }

        ammoArray.sort((a,b) => {
            if (a.penetration > b.penetration) return 1
            if (a.penetration < b.penetration) return -1
            return 0
        })

        for (const ammo of ammoArray) {
            if (ammo.damage > 200) {
                trace.x.push(200)
            } else {
                trace.x.push(ammo.damage)
            }

            trace.y.push(ammo.penetration)

            if (ammo.price > 5000) {
                trace.z.push(5000)
            } else {
                trace.z.push(ammo.price)
            }
            trace.text.push(ammo.name)
        }

        if (caliber === "Caliber12g" || 
            caliber === "Caliber9x18PM" ||
            caliber === "Caliber762x25TT" ||
            caliber === "Caliber9x21" ||
            caliber === "Caliber20g"
        ) {
            trace.visible = 'legendonly'
        }
        traces.push(trace)
    }

    return traces
}

export function createTraces2D(ammoData) {
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

        let ammoArray = []

        for (const [_, ammo] of Object.entries(ammoData.record[caliber])) {
            ammoArray.push(ammo)
        }

        ammoArray.sort((a,b) => {
            if (a.penetration > b.penetration) return 1
            if (a.penetration < b.penetration) return -1
            return 0
        })

        for (const ammo of ammoArray) {
            if (ammo.damage > 200) {
                trace.x.push(200)
            } else {
                trace.x.push(ammo.damage)
            }

            trace.y.push(ammo.penetration)

            if (ammo.price > 5000) {
                trace.z.push(5000)
            } else {
                trace.z.push(ammo.price)
            }
            trace.text.push(ammo.name)
        }

        if (caliber === "Caliber12g" || 
            caliber === "Caliber9x18PM" ||
            caliber === "Caliber762x25TT" ||
            caliber === "Caliber9x21" ||
            caliber === "Caliber20g"
        ) {
            trace.visible = 'legendonly'
        }
        traces.push(trace)
    }

    return traces
}