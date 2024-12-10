defmodule Day3 do
  @moduledoc """
  Documentation for `Day3`.
  """

  @doc """
  Hello world.

  ## Examples

      iex> Day3.hello()
      :world

  """
  def hello do
    :world
  end

  # read fro the file
  #
  # then pass the regix
  #
  # get the data back
  #
  # do the operations that are in the regix
    #
  @spec readfile() :: String.t()
  def readfile do
    {:ok, filedata} = File.read("input.txt")
    filedata
    |> regix()
  end

  def regix(filedata) do


    data  = Regex.scan(~r/mul\(\d{1,3}\,\d{1,3}\)/, filedata)
    Enum.map(data, fn [x] -> 
      [x, y] =  
         Regex.scan(~r/\d+/, x)
          |> List.flatten()
      String.to_integer(x) * String.to_integer(y)
      end)
      |> sum()
  end

    @spec sum(list()) :: integer()
    def sum([head | tail]) do
      head + sum(tail)
    end

    def sum([]) do
      0
    end
end
