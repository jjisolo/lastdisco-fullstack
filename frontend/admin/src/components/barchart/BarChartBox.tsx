import './BarChartBox.scss'
import {BarChart, Bar, ResponsiveContainer, Tooltip} from 'recharts';

type Props = {
    title: string;
    color: string;
    dataKey: string;
    chartData: object[];
}

const BarChartBox = (props: Props) => {
    return (
        <div className="barChart">
            <h1 className="title">
                { props.title }
            </h1>
            <div className="chart">
                <ResponsiveContainer width="99%" height={100}>
                    <BarChart width={150} height={40} data={props.chartData}>
                        <Tooltip
                            contentStyle={{ background: "#2a3447", border: "5px"}}
                            labelStyle={{ display: "none"}}
                            cursor={{ fill: "none"}}
                        />
                        <Bar dataKey={props.dataKey} fill={props.color}/>
                    </BarChart>
                </ResponsiveContainer>
            </div>
        </div>
    )
}

export default BarChartBox