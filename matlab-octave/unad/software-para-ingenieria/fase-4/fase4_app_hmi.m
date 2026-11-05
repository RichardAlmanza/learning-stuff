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
        GeneraciontotalPanel            matlab.ui.container.Panel
        LabelTotalBalance               matlab.ui.control.Label
        BalancePanel                    matlab.ui.container.Panel
        LabelBalance                    matlab.ui.control.Label
        UITable                         matlab.ui.control.Table
        GeneracionvsDemandaPanel        matlab.ui.container.Panel
        UIAxes                          matlab.ui.control.UIAxes
        GridLayoutTopSection            matlab.ui.container.GridLayout
        EstadodelSistemaPanel           matlab.ui.container.Panel
        GridLayoutSystemStatePanel      matlab.ui.container.GridLayout
        GridLayoutWaterPump             matlab.ui.container.GridLayout
        BombaLabel                      matlab.ui.control.Label
        WaterPumpLamp                   matlab.ui.control.Lamp
        GridLayoutPublicLightingState   matlab.ui.container.GridLayout
        AlumbradoLabel                  matlab.ui.control.Label
        PublicLightingLamp              matlab.ui.control.Lamp
        GridLayoutSchoolState           matlab.ui.container.GridLayout
        EscuelaLabel                    matlab.ui.control.Label
        SchoolLamp                      matlab.ui.control.Lamp
        BateriaGauge                    matlab.ui.control.SemicircularGauge
        BateriaGaugeLabel               matlab.ui.control.Label
        ControldelSimuladorPanel        matlab.ui.container.Panel
        GridLayoutSimulationControlPanel  matlab.ui.container.GridLayout
        ReproducirButton                matlab.ui.control.StateButton
        SliderTime                      matlab.ui.control.Slider
        ReiniciarButton                 matlab.ui.control.Button
        VariablesdeEntradaPanel         matlab.ui.container.Panel
        GridLayoutInputVariablePanel    matlab.ui.container.GridLayout
        GridLayoutSeedControl           matlab.ui.container.GridLayout
        EditField                       matlab.ui.control.NumericEditField
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
        ScaleDemand = 1
        ScaleWind = 1
        ScaleSolar = 1
        DemandRecords
        SolarRecords
        WindRecords
        
    end
    

    % Callbacks that handle component events
    methods

        % Code that executes after component creation
        function startupFcn(app)
            [a, b] = generatePowerGenerationData;
            app.SolarRecords = [];
            app.SliderSolar.Value = max(app.SolarRecords);
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