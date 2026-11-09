classdef fase4_app_hmi < matlab.apps.App

    % Used to locate and load the app's XML configuration file
    properties (Access = public, Constant)
        AppConfigFilename = './fase4_app_hmi.xml'; % File path to the app configuration file containing component layout and settings
    end

    % Properties that correspond to app components
    properties (Access = public)
        UIFigure                        matlab.ui.Figure
        GridLayout                      matlab.ui.container.GridLayout
        GridLayoutTitle                 matlab.ui.container.GridLayout
        Image                           matlab.ui.control.Image
        AppTitle                        matlab.ui.control.Label
        GridLayoutBottomSection         matlab.ui.container.GridLayout
        RegistroPanel                   matlab.ui.container.Panel
        GridLayout7                     matlab.ui.container.GridLayout
        BatteryStateLabel               matlab.ui.control.Label
        GridLayout8                     matlab.ui.container.GridLayout
        BalanceTotalPanel               matlab.ui.container.Panel
        GridLayoutBalance_2             matlab.ui.container.GridLayout
        LabelTotalBalance               matlab.ui.control.Label
        kWLabel_6                       matlab.ui.control.Label
        BalancePanel                    matlab.ui.container.Panel
        GridLayoutBalance               matlab.ui.container.GridLayout
        LabelBalance                    matlab.ui.control.Label
        kWLabel_5                       matlab.ui.control.Label
        UITable                         matlab.ui.control.Table
        GeneracionvsDemandaPanel        matlab.ui.container.Panel
        UIAxes                          matlab.ui.control.UIAxes
        GridLayoutTopSection            matlab.ui.container.GridLayout
        EstadodelSistemaPanel           matlab.ui.container.Panel
        GridLayoutSystemStatePanel      matlab.ui.container.GridLayout
        GridLayoutWaterPump             matlab.ui.container.GridLayout
        BombaLabel                      matlab.ui.control.Label
        WaterPumpLamp                   matlab.ui.control.Lamp
        GridLayoutBatteryLabel          matlab.ui.container.GridLayout
        PercentageLabel                 matlab.ui.control.Label
        LabelBattery                    matlab.ui.control.Label
        LabelBatteryValue               matlab.ui.control.Label
        GridLayoutPublicLightingState   matlab.ui.container.GridLayout
        AlumbradoLabel                  matlab.ui.control.Label
        PublicLightingLamp              matlab.ui.control.Lamp
        GridLayoutSchoolState           matlab.ui.container.GridLayout
        EscuelaLabel                    matlab.ui.control.Label
        SchoolLamp                      matlab.ui.control.Lamp
        GaugeBattery                    matlab.ui.control.SemicircularGauge
        ControldelSimuladorPanel        matlab.ui.container.Panel
        GridLayoutSimulationControlPanel  matlab.ui.container.GridLayout
        SliderTime                      matlab.ui.control.Slider
        ButtonResetSimulation           matlab.ui.control.Button
        ButtonRunSimulation             matlab.ui.control.StateButton
        VariablesdeEntradaPanel         matlab.ui.container.Panel
        GridLayoutInputVariablePanel    matlab.ui.container.GridLayout
        GridLayoutSeedControl           matlab.ui.container.GridLayout
        EditFieldSeed                   matlab.ui.control.NumericEditField
        GenerarsemillaButton            matlab.ui.control.Button
        GridLayoutDemandHeader          matlab.ui.container.GridLayout
        kWLabel_3                       matlab.ui.control.Label
        DemandValueLabel                matlab.ui.control.Label
        DemandadelaComunidadLabel       matlab.ui.control.Label
        GridLayoutWindHeader            matlab.ui.container.GridLayout
        kWLabel_4                       matlab.ui.control.Label
        WindValueLabel                  matlab.ui.control.Label
        EnergiaEolicaLabel              matlab.ui.control.Label
        GridLayoutSolarHeader           matlab.ui.container.GridLayout
        kWLabel                         matlab.ui.control.Label
        SolarValueLabel                 matlab.ui.control.Label
        EnergiaSolarLabel               matlab.ui.control.Label
        SliderDemand                    matlab.ui.control.Slider
        SliderWind                      matlab.ui.control.Slider
        SliderSolar                     matlab.ui.control.Slider
    end

    
    properties (Access = private)
        GrayColor = [0.30, 0.30, 0.30]
        GreeColor = [0, 1, 0]
        DemandRecords
        SolarRecords
        WindRecords
        DeltaRecordsScaled
        SocRecordsScaled
        StatesRecords
        Soc
        Sm
        LampColorsArray
        Hours
        CursorLine
        DemandPlot
        SolarPlot
        WindPlot
        DemandMax
        SolarMax
        WindMax
        Timer1S
        
    end
    
    methods (Access = private)
        
        function [YData, valueText] = UpdatePlot(~, value, maxRecord, records)
            scalePlot = value / maxRecord;
            % Update plot
            YData = records .* scalePlot;
            % Update value label
            valueText = num2str(value);
        end
        
        function UpdateRecords(app, seed)
            % Update records
            [app.SolarRecords, app.WindRecords] = generatePowerGenerationData(seed);
            app.DemandRecords = generateCommunityDemandData(seed);

            % Update Max values
            app.SolarMax = max(app.SolarRecords);
            app.WindMax = max(app.WindRecords);
            app.DemandMax = max(app.DemandRecords);
        end
        
        function UpdateHour(~, slider, callbkfnc, min, max)
            newValue = slider.Value + 1;
            if newValue > max
                newValue = min;
            end

            slider.Value = newValue;
            callbkfnc(slider);

        end
        
        function UpdateDeltaSocAndStatesRecords(app)
            app.DeltaRecordsScaled = app.SolarPlot.YData + ...
                app.WindPlot.YData - app.DemandPlot.YData;
            [~, app.SocRecordsScaled] = app.Soc.ChargeRecords(app.DeltaRecordsScaled);
            [~, app.StatesRecords] = app.Sm.EvaluateRecords(app.SocRecordsScaled);
            app.LabelTotalBalance.Text = num2str(sum(app.DeltaRecordsScaled), 5);

        end

        function UpdateTable(app)
            totalGen = app.SolarPlot.YData + app.WindPlot.YData;
            newTable = table((1:24)', totalGen', app.DemandPlot.YData', ...
                app.SocRecordsScaled');
            app.UITable.Data = newTable;
        end

        function UpdateLamps(app, hour)
            len = length(app.LampColorsArray);

            for index = 1:len
                lamp = app.LampColorsArray(index);
                state = app.StatesRecords(hour, index);

                if state
                    lamp.Color = app.GreeColor;
                else
                    lamp.Color = app.GrayColor;
                end
            end
        end

        function UpdateBatteryGauge(app, hour)
            chargePercentage = app.SocRecordsScaled(hour);
            app.GaugeBattery.Value = chargePercentage;
            app.LabelBatteryValue.Text = num2str(chargePercentage, 3);
        end

        function UpdateSystemStateUIPanel(app, hour)
            if isstring(hour)
                if hour == "current"
                    hour = app.SliderTime.Value;
                end
            end

            app.UpdateLamps(hour);
            app.UpdateBatteryGauge(hour);
            app.LabelBalance.Text = num2str(app.DeltaRecordsScaled(hour), 3);
        end
    end
    

    % Callbacks that handle component events
    methods

        % Code that executes after component creation
        function startupFcn(app)
            % Load default data
            app.Hours = 1:24;
            seed = 0;
            app.EditFieldSeed.Value = seed;
            UpdateRecords(app, seed);
            % Set Sliders information
            app.SliderSolar.Value = app.SolarMax;
            app.SliderWind.Value = app.WindMax;
            app.SliderDemand.Value = app.DemandMax;
            % Set text value of sliders
            app.SolarValueLabel.Text = num2str(app.SolarMax);
            app.WindValueLabel.Text = num2str(app.WindMax);
            app.DemandValueLabel.Text = num2str(app.DemandMax);

            % Load SOC
            app.Soc = StateOfCharge(100,50,0,100);
            app.GaugeBattery.Value = app.Soc.State;

            % Load StateManager
            app.Sm = StateManager( ...
                ["School", "WaterPump", "PublicLighting"], ...
                [0, 20, 40] ...
                );
            app.Sm = app.Sm.Evaluate(app.Soc.State);

            % Set Lamps, it is important to keep the same order as defined
            % at StateManager construction
            app.LampColorsArray = [
                app.SchoolLamp, ...
                app.WaterPumpLamp, ...
                app.PublicLightingLamp
                ];
            
            % Set chart
            hold(app.UIAxes, "on");
            grid(app.UIAxes, "on");
            xticks(app.UIAxes, app.Hours);
            xlim(app.UIAxes, [1,24]);

            app.SolarPlot = plot(app.UIAxes, ...
                app.Hours, app.SolarRecords, ...
                "DisplayName", "Solar", ...
                "Color", "#FFB300");
            app.WindPlot = plot(app.UIAxes, ...
                app.Hours, app.WindRecords, ...
                "DisplayName", "Eolica", ...
                "Color", "#4DA3FF");
            app.DemandPlot = plot(app.UIAxes, ...
                app.Hours, app.DemandRecords, ...
                "DisplayName", "Demanda", ...
                "Color", "#FF6B52");

            hold(app.UIAxes, "off");
            legend(app.UIAxes);

            % Update delta and soc records, they depend on the YData of
            % plot objects
            app.UpdateDeltaSocAndStatesRecords();
            app.UpdateTable();
            app.UpdateSystemStateUIPanel(1);
            app.UITable.SelectionType = "row";
            app.UITable.Selection = 1;
            app.LabelBalance.Text = num2str(app.DeltaRecordsScaled(1), 3);

            % Set cursor
            app.CursorLine = xline(app.UIAxes, 1, ...
                "Color", "r", ...
                "LineWidth", 2, ...
                "LineStyle", "--", ...
                "DisplayName", "Hora");

            % Set timer
            app.Timer1S = timer( ...
                "StartDelay", 1, ...
                "Period", 1, ...
                "ExecutionMode","fixedRate" ...
                );
            app.Timer1S.TimerFcn = @(~, ~) app.UpdateHour(app.SliderTime, @app.SliderTimeValueChanging, 1, 24);

        end

        % Value changing function: SliderSolar
        function SliderSolarValueChanging(app, event)
            [app.SolarPlot.YData, app.SolarValueLabel.Text] = UpdatePlot(app, event.Value, app.SolarMax, app.SolarRecords);
            app.UpdateDeltaSocAndStatesRecords();
            app.UpdateTable();
            app.UpdateSystemStateUIPanel("current");
        end

        % Value changing function: SliderWind
        function SliderWindValueChanging(app, event)
            [app.WindPlot.YData, app.WindValueLabel.Text] = UpdatePlot(app, event.Value, app.WindMax, app.WindRecords);
            app.UpdateDeltaSocAndStatesRecords();
            app.UpdateTable();
            app.UpdateSystemStateUIPanel("current");
        end

        % Value changing function: SliderDemand
        function SliderDemandValueChanging(app, event)
            [app.DemandPlot.YData, app.DemandValueLabel.Text] = UpdatePlot(app, event.Value, app.DemandMax, app.DemandRecords);
            app.UpdateDeltaSocAndStatesRecords();
            app.UpdateTable();
            app.UpdateSystemStateUIPanel("current");
        end

        % Value changed function: EditFieldSeed
        function EditFieldSeedValueChanged(app, event)
            UpdateRecords(app, event.Value);

            % Update sliders and plot
            app.SliderSolarValueChanging(app.SliderSolar);
            app.SliderWindValueChanging(app.SliderWind);
            app.SliderDemandValueChanging(app.SliderDemand);
        end

        % Button pushed function: GenerarsemillaButton
        function GenerarsemillaButtonPushed(app, event)
            rng("shuffle");
            % 2^31 = 2147483648
            app.EditFieldSeed.Value = randi([0, 2147483648], 1, 1, "double");

            app.EditFieldSeedValueChanged(app.EditFieldSeed);
        end

        % Value changing function: SliderTime
        function SliderTimeValueChanging(app, event)
            hour = event.Value;
            app.CursorLine.Value = hour;

            app.UpdateSystemStateUIPanel(hour);

            scroll(app.UITable, "row", hour);
            app.UITable.Selection = hour;
        end

        % Value changed function: ButtonRunSimulation
        function ButtonRunSimulationValueChanged(app, event)
            if event.Value
                start(app.Timer1S);
            else
                stop(app.Timer1S);
            end

        end

        % Button pushed function: ButtonResetSimulation
        function ButtonResetSimulationPushed(app, event)
            app.SliderTime.Value = 1;
            app.SliderTimeValueChanging(app.SliderTime);
        end

        % Selection changed function: UITable
        function UITableSelectionChanged(app, event)
            selection = app.UITable.Selection;
            app.SliderTime.Value = selection;
            app.SliderTimeValueChanging(app.SliderTime);
        end
    end

    % App creation
    methods (Access = public)

        % Construct app
        function app = fase4_app_hmi(varargin)
            app = app@matlab.apps.App(varargin{:});

            if nargout == 0
                clear app
            end
        end
    end
end