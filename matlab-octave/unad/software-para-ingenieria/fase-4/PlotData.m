classdef PlotData
  properties
    Data (1,:) double
    LineWidth (1,1) double
    LineSpec
    Legend
  end

  methods
    function pData = PlotData(data, lineSpec, lineWidth, displayName)
      pData.Data = data;
      pData.LineSpec = lineSpec;
      pData.LineWidth = lineWidth;
      pData.Legend = displayName;
    end
  end
end

