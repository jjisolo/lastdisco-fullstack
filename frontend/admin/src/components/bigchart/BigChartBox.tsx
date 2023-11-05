import './BigChartBox.scss'
import {Area, AreaChart , ResponsiveContainer, Tooltip, XAxis, YAxis} from "recharts";

const data = [
    {
        name: "Sun",
        pets: 4000,
        games: 2400,
        pictures: 2400,
    },
    {
        name: "Mon",
        pets: 3000,
        games: 1398,
        pictures: 2210,
    },
    {
        name: "Tue",
        pets: 2000,
        games: 9800,
        pictures: 2290,
    },
    {
        name: "Wed",
        pets: 2780,
        games: 3908,
        pictures: 2000,
    },
    {
        name: "Thu",
        pets: 1890,
        games: 4800,
        pictures: 2181,
    },
    {
        name: "Fri",
        pets: 2390,
        games: 3800,
        pictures: 2500,
    },
    {
        name: "Sat",
        pets: 3490,
        games: 4300,
        pictures: 2100,
    },
];

const BigChartBox = () => {

    return (
        <div className="bigChartBox">
            <h1 style={{marginBottom: "14px"}}>
                Аналитика продаж
            </h1>

            <div className="chart">
                <ResponsiveContainer width="99%" height="100%">
                    <AreaChart
                        data={data}
                        margin={{
                            top: 10,
                            right: 30,
                            left: 0,
                            bottom: 0,
                        }}
                    >
                        <XAxis dataKey="name" />
                        <YAxis />
                        <Tooltip />
                        <Area type="monotone" dataKey="pets" stackId="1" stroke="#8884d8" fill="#8884d8" />
                        <Area type="monotone" dataKey="games" stackId="1" stroke="#82ca9d" fill="#82ca9d" />
                        <Area type="monotone" dataKey="pictures" stackId="1" stroke="#ffc658" fill="#ffc658" />
                    </AreaChart>
                </ResponsiveContainer>
            </div>
        </div>

    )
}

export default BigChartBox