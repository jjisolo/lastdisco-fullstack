import './Home.scss'
import TopBox from "../../components/topbox/TopBox.tsx";
import ChartBox from "../../components/chartbox/ChartBox.tsx";
import {ChartConversion, ChartProduct, ChartRevenue, ChartUsers, ChartBoxVisit, ChartBoxRevenue} from "./Template.tsx";
import BarChartBox from "../../components/barchart/BarChartBox.tsx";
import PieChartBox from "../../components/piechart/PieChartBox.tsx";
import BigChartBox from "../../components/bigchart/BigChartBox.tsx";

const Home = () => {
    return (
        <div className="home">
            <div className="box box1">
                <TopBox/>
            </div>
            <div className="box box2">
                <ChartBox {...ChartUsers}/>
            </div>
            <div className="box box3">
                <ChartBox {...ChartProduct}/>
            </div>
            <div className="box box4">
                <PieChartBox />
            </div>
            <div className="box box5">
                <ChartBox {...ChartConversion}/>
            </div>
            <div className="box box6">
                <ChartBox {...ChartRevenue}/>
            </div>
            <div className="box box7">
                <BigChartBox />
            </div>
            <div className="box box8">
                <BarChartBox {...ChartBoxVisit} />
            </div>
            <div className="box box9">
                <BarChartBox {...ChartBoxRevenue} />
            </div>
        </div>
    )
}

export default Home